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

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintln(employee Employee) {
	/*
		1. 建立连接 tcp http
		2. 将 employee 对象数列化 json 字符串 序列化 发送 json 字符串
		3. 发送成功后 实际上接受到的是一个二进制数据
		4. 等待服务器发送结果 将服务器返回的数据解析成 printresult 对象 反序列化
	*/

	//http 2.0 微服务 rpc 接口一般需要长连接 而且需要 性能 一般 用 http 2.0 或者自己封装一套 tcp 协议
	// 还有传输数据协议 xml  json protobuf msgpack
}

func Add(a, b int) int {
	total := a + b
	return total
}

// rpc要素分析

// rpc 技术架构设计上有四部分组成 客户端 客户端存根 服务端 服务端存根

// 客户端 服务调用发起方 服务消费者

// 客户端存根 指程序运行在客户端所在的机器上 主要用来存储调用的服务器地址 改程序
// 还负责将客户端请求远端服务器程序的数据信息打包成数据包 通过网络发送给服务端 接受服务端发送的数据包 返回给客户端

// 服务端 远端计算机机器运行的程序 其中有客户端要调用的方法
// 服务端存根 接收客户 stub 程序通过网络发送的请求消息数据包 调用服务端功能方法 完成调用
