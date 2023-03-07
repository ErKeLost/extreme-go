package main

import (
	"fmt"
	"time"
)

// 并发编程

// python java php 多线程 多进程

// 内存 线程切换 用户级线程 绿程 协程

// 携程 go语言的协程 go语言没有线程 -goroutine

func asyncPrint() {

	for {
		time.Sleep(time.Second)
		fmt.Println("kobe")
	}
}

// 主协程
func main() {
	// 主死 异步没来得及跑
	//go asyncPrint() // 生成协程
	//
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		fmt.Println("kobe")
	//	}
	//}()
	//
	//for i := 0; i < 100; i++ {
	//	go func(i int) {
	//		time.Sleep(time.Second)
	//		fmt.Println(i)
	//	}(i) // 值传递 就会复制 就能拿到变量了
	//}
	//fmt.Println("主协程")
	//time.Sleep(10 * time.Second)
	bb := &aa{888}
	//fmt.Println(bb)
	fmt.Printf("%p\r\n", bb)
	Super(bb)
	fmt.Println(bb.age)

}

//func main() {
//	bb := &aa{"adny", 888}
//	Super(bb)
//	fmt.Println(bb.age)
//}

func Super(s aa) {
	fmt.Println(s)
	//fmt.Printf("%p\r\n", s)
	s = &aa{666}
	//s.age = 6666
	//fmt.Println(s, 666)
	//fmt.Printf("%p\r\n", s)

}

type aa struct {
	age int
}
