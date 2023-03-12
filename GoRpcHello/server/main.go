package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply 的值
	*reply = "hello," + request
	return nil
}

func main() {
	// rpc 三步走策略

	// 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")

	// 注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 启动服务
	conn, _ := listener.Accept() // 当一个新的链接进来的时候
	for {
		rpc.ServeConn(conn)
	}
}
