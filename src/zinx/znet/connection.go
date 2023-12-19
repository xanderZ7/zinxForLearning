package znet

import (
	"fmt"
	"net"
	"zinx/src/zinx/ziface"
)

// 连接模块
type Connection struct {
	// 当前连接的 socket tcp套接字
	Conn *net.TCPConn

	// 连接的 ID
	ConnID uint32

	// 当前的连接状态
	isClosed bool

	// 当前连接所绑定的处理业务方法 API
	handleAPI ziface.HandleFunc

	// 告知当前连接退出的 channel
	ExitChan chan bool

	// 该链接处理的方法
	Router ziface.IRouter
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {

	c := &Connection{
		Conn:     conn,
		ConnID:   connId,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}

	return c
}

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")

	defer fmt.Println("connId = ", c.ConnID, "Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中，最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		// 得到当前conn数据的Request请求数据
		req := Request{conn: c, data: buf}

		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		// 从路由中，找到注册绑定的conn对应的router调用

		// // 调用当前连接所绑定的 HandleAPI
		// if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		// 	fmt.Println("ConnID = ", c.ConnID, "handle is error", err)
		// 	break
		// }
	}
}

// 启动连接，让当前的连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)

	// 启动从当前连接的读数据业务
	go c.StartReader()
}

// 停止连接，结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true

	// 关闭socket连接
	c.Conn.Close()

	// 资源回收
	close(c.ExitChan)
}

// 获取当前连接的绑定 socket conn
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接模块的连接 ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的 tcp状态 ip port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
