package main

import "zinx/src/zinx/znet"

func main() {

	// 1. 创建一个 server 句柄，使用 Zinx 的 api
	s := znet.NewServer("[zinx V0.2]")

	// 2. 启动 server
	s.Serve()
}
