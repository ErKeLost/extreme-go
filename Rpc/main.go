package main

import "fmt"

type Employee struct {
	Name    string
	Address string
}

func main() {

	// 什么是rpc
	/*
		rpc remote procedure call 远程过程调用 简单的理解就是一个节点请求另一个节点的服务

		对应rpc的是本地过程调用 函数调用是最常见的本地过程调用

		将本地过程调用编程远程调用会面临各种问题
	*/
	// 现在把 add 函数想象成一个远程函数嗲用 电商系统 有一套库存系统 现在你需要调用库存系统的服务 如何调用

	// 函数调用参数 如何传递 json json 是一种数据格式协议 xmd protobuf msgpack

	//
	fmt.Println(Add(33, 66))

	// 把这个打印结果传递到另一台服务器上
	// 这样不可行 可行的方式就是将struct 转成json
	fmt.Println(Employee{
		Name:    "erkelost",
		Address: "beijing",
	})
}

func Add(a, b int) int {
	total := a + b
	return total
}
