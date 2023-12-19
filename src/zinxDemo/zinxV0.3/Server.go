package main

import (
	"fmt"
	"zinx/src/zinx/ziface"
	"zinx/src/zinx/znet"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) PreHandle(request ziface.IRequest) {

	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}

}

func (this *PingRouter) Handle(request ziface.IRequest) {

	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}

}

func (this *PingRouter) PostHandle(request ziface.IRequest) {

	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("After ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}

}

func main() {

	// 1. 创建一个 server 句柄，使用 Zinx 的 api
	s := znet.NewServer("[zinx V0.3]")

	// 给当前zinx框架添加一个自定义router
	s.AddRouter(&PingRouter{})

	// 2. 启动 server
	s.Serve()
}
