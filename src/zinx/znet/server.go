package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/src/zinx/ziface"
)

// iServer 的接口实现，定义一个Server的服务器模块
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的ip
	IP string
	// 服务器监听的端口
	Port int
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] Call back to client")

	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("Call back to client error")
	}

	return nil
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at Ip: %s, Port %d is Starting\n", s.IP, s.Port)

	go func() {
		// 1. 获取一个 TCP 的 Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		// 2. 监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("Start Zinx server suncc,", s.Name, "succ, Listenning...")
		var cid uint32
		cid = 0

		// 3. 阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			// 如果有客户端连接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 将处理新连接的业务方法 和 conn 进行绑定，得到我们的连接模块
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			// 启动当前的连接业务处理
			go dealConn.Start()
			// // 已经与客户端建立连接，做一个最大512字节长度的回显业务
			// go func() {
			// 	for {
			// 		buf := make([]byte, 512)
			// 		cnt, err := conn.Read(buf)
			// 		if err != nil {
			// 			fmt.Println("recv buf err", err)
			// 			continue
			// 		}

			// 		fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)

			// 		// 回显业务
			// 		if _, err := conn.Write(buf[:cnt]); err != nil {
			// 			fmt.Println("write back buf err", err)
			// 			continue
			// 		}
			// 	}
			// }()

		}
	}()
}

// 停止服务器
func (s *Server) Stop() {

}

// 运行服务器
func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO

	// 进入阻塞状态
	select {}
}

// 初始化 Server 模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}
