package main

func main() {
	// java 交给虚拟机来进行交互 jvm

	// 比如说多个线程 调度的话 要在操作系统运行 导致线程内存占用高

	// 函数级别的切换 协程

	// go 语言的 调度 gmp 调度  申请一些线程

	// goroutine  x  10000 映射在一个线程上

	// 如何映射 生成一个中间调度器  M : N

	// m个线程 对应 n 个 goroutine

	// GMP  goroutine -> M- thread 线程
}

//
//import (
//	"fmt"
//	"time"
//)
//
//// 并发编程
//
//// python java php 多线程 多进程
//
//// 内存 线程切换 用户级线程 绿程 协程
//
//// 携程 go语言的协程 go语言没有线程 -goroutine
//
//func asyncPrint() {
//
//	for {
//		time.Sleep(time.Second)
//		fmt.Println("kobe")
//	}
//}
//
//// 主协程
//func main() {
//	// 主死 异步没来得及跑
//	//go asyncPrint() // 生成协程
//	//
//	//go func() {
//	//	for {
//	//		time.Sleep(time.Second)
//	//		fmt.Println("kobe")
//	//	}
//	//}()
//	//
//	//for i := 0; i < 100; i++ {
//	//	go func(i int) {
//	//		time.Sleep(time.Second)
//	//		fmt.Println(i)
//	//	}(i) // 值传递 就会复制 就能拿到变量了
//	//}
//	//fmt.Println("主协程")
//	//time.Sleep(10 * time.Second)
//	bb := &aa{888}
//	//fmt.Println(bb)
//	fmt.Printf("%p\r\n", bb)
//	//Super(bb)
//	fmt.Println(bb.age)
//
//}
//
////func main() {
////	bb := &aa{"adny", 888}
////	Super(bb)
////	fmt.Println(bb.age)
////}
//
//func Super(s aa) {
//	fmt.Println(s)
//	//fmt.Printf("%p\r\n", s)
//	s = &aa{666}
//	//s.age = 6666
//	//fmt.Println(s, 666)
//	//fmt.Printf("%p\r\n", s)
//
//}
//
//type aa struct {
//	age int
//}
