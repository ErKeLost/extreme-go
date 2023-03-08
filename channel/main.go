package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func main() {
	/*
		不要通过共享内存来通信 要通信来实现内存共享
		多线程编程的时候 两个 goroutine 最常用的方式 通过一个全局变量

		然后修改的时候 玩锁

		使用 channel 更加简单 通过一个 goroutine 到 另一个 goroutine 进行通信
	*/

	//var msg chan int
	//
	//// 有缓冲和无缓冲的channel
	//msg = make(chan int, 2) // channel 初始化值 如果是0 就会被阻塞 下面代码永远运行不到
	//
	//fmt.Println(msg)
	//// 语法糖传值
	////msg <- 1 // 存放值
	////msg <- 1 // 存放值
	////
	////data := <-msg // 取值
	//
	//// 默认值是nil
	//if msg == nil {
	//	fmt.Println("msg is nil")
	//}
	//
	//// 有缓冲channel 和 无缓冲 channel 应用场景
	//
	//// 无缓冲适用于 通知 b要第一时间知道a已经完成
	//
	//// 有缓存 channel 用于消费者和生产者之间的通信
	//
	////channel 应用场景
	//
	////forrange语法
	//
	//fmt.Println("--------------FORRANGE")
	//
	//go func(msg chan int) {
	//	for data := range msg {
	//		fmt.Println(data)
	//	}
	//
	//	fmt.Println("all down")
	//}(msg)
	//
	//msg <- 1
	//msg <- 2
	//
	//msg <- 999
	//close(msg) // 已经关闭的channel 可以继续取值 但是不能继续赋值了
	//
	//time.Sleep(time.Second * 10)
	//
	//// 单项channel 的应用场景
	//
	//// 默认情况下 channel 是双向的 既可以接受也能发送
	//
	//// 但是经常 channel 作为了参数进行了从传递 就是 vue 单向数据流
	//
	//var ch1 chan int // 双向 channel
	//
	//var ch2 chan<- float64 // 单向channel 只能写float64的数据
	//
	//var ch3 <-chan int // 只能读取
	//
	//c := make(chan int, 3)
	//
	//var send chan<- int = c // 只能写
	//
	//var read <-chan int = c // 只能读
	//
	//send <- 1
	////<-send 报错
	//<-read

	c := make(chan int)

	go producer(c)

	go consumer(c)

	//time.Sleep(10 * time.Second)

	// channel 实现交叉打印 面试题

	/*
		使用两个 goroutine 交叉打印序列  一个打印数字一个打印字母
	*/

	go printNum()
	go printLetter()

	number <- true

	time.Sleep(10 * time.Second)

}

func printNum() {
	i := 1
	for {
		<-number
		// 等待另一个 goroutine 改打印了
		fmt.Printf("%d%d", i, i+1)
		i += 2

		letter <- true
	}
}

func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		<-letter
		// 等待另一个 goroutine 改打印了
		if i >= len(str) {
			return
		}
		fmt.Print(str[i : i+2])
		i += 2
		number <- true
	}
}

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num=%d\r\n", num)
	}
}
