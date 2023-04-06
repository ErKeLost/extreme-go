package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立链接

	client, err := rpc.Dial("tcp", ":1234")

	var reply *string
	err = client.Call("HelloService.Hello", "erkelost", &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println(*reply)
	// 调用方法
}
