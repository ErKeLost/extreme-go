package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 实例化一个server
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		fmt.Println(jsonrpc.NewServerCodec(conn))
		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		return
	}

	// 注册处理逻辑 handler

	// 启动服务
}
