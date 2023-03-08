package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享全局变量的方式
//var done bool

// 全局channel 的方式
// var doneBool chan bool
var done = make(chan struct{}) // 空结构体 // chanel 要初始化

var lock sync.Mutex

func g1(ch chan struct{}) {
	time.Sleep(time.Second)
	lock.Lock()
	defer lock.Unlock()
	//done = true
	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	time.Sleep(time.Second * 2)
	lock.Lock()
	defer lock.Unlock()
	//done = true
	ch <- struct{}{}
}

func main() {
	// select 类似与 switch case 语句  但是 select 功能 跟linux io select poll epoll
	// select 主要作用于 多个 channel

	//从代码中 选择一个已经就绪的channel

	// 现在有两个 goroutine 都在执行

	//我在主的goroutine 当某一个执行完成之后 立马知道哪一个完成了
	g1Channel := make(chan struct{}) // 初始化值为0 就会阻塞
	g2Channel := make(chan struct{}) // 初始化值为0 就会阻塞
	// 放值进去就会阻塞
	msg := make(chan string, 0)
	//msg <- "666" // 这段代码 永远阻塞 就永远走不到  //无缓冲
	// 无缓冲的 channel 先写 先读 都不行
	//data := <-msg

	// 解决无缓冲的问题 就接受一个 goroutine

	go func(msg chan string) {
		// TODO
		// 拿到无缓冲的msg
		data := <-msg
		fmt.Println(data)

		// 为什么能行呢 go 语言中有一种 happen-before 的机制 有一个保障

		//
	}(msg)

	msg <- "erkelpst"
	// 无缓冲就会阻塞
	go g1(g1Channel)

	go g2(g2Channel)
	//for {
	//	if done {
	//		fmt.Println("done")
	//		time.Sleep(10 * time.Millisecond)
	//		return
	//	}
	//}
	//<- done

	// 某一个分支就绪了就执行该分支 如果两个分支都就绪了 随机执行哪一个
	select {
	case <-g1Channel:
		fmt.Println("g1Channel 先执行")
	case <-g2Channel:
		fmt.Println("g2Channel 先执行")

	}
}
