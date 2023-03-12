package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立连接

	client, _ := rpc.Dial("tcp", "localhost:1234")
	// new 内存分配一个空间
	var reply *string = new(string)

	// var reply string string 默认值 传递的时候把他编程指针
	_ = client.Call("HelloService.Hello", "我是", reply)
	fmt.Println(*reply)

	// 跨语言调用 go语言 的 rpc 序列化 和 反序列化协议是什么
	// 能否替换成常见的序列化
}
