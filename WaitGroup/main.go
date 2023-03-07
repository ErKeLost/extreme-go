package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//func main() {
//	// 子 goroutine 如何通知 主的 goroutine
//	// 主 goroutine 如何 知道 子的goroutine 已经结束了
//	var wg sync.WaitGroup
//
//	// 监控多少个 goroutine 结束
//	//wg.Add(100)
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go func(i int) {
//			// 函数执行完成调用一次 循环一次调用一次
//			defer wg.Done()
//			fmt.Println(i)
//			//wg.Done()
//		}(i)
//	}
//	// 等到 goroutine 全部执行完 就结束
//	wg.Wait()
//
//	// wait group 主要用于 goroutine 执行等待 add 方法 要和 done 方法配套
//	fmt.Println("all dont wait group")
//
//	// 如何使用 锁
//
//	fmt.Println("lock ----------------")
//
//	/*
//		锁 - 资源竞争
//	*/
//
//}

var total int32
var wg sync.WaitGroup
var lock sync.Mutex

// var lock2 := lock  // 锁 不能复制 本质 struct 值传递 状态被修改
func add() {
	defer wg.Done()
	for i := 0; i < 10000000; i++ {
		atomic.AddInt32(&total, 1)
		//lock.Lock()
		//total += 1
		//lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 10000000; i++ {
		atomic.AddInt32(&total, -1)
		//lock.Lock()
		//total -= 1
		//lock.Unlock()
	}
}

func main() {

	// 互斥锁
	//wg.Add(2)
	//go add()
	//go sub()
	//wg.Wait()
	//fmt.Println(total)

	// 不知道到底谁会被先执行 赋值 调度 保证原子性的执行

	// 读写锁 锁本质上是将并行的代码串行化 lock 会影响性能 并发编程避免使用锁

	// 即使是设计锁 那么也应该尽量的保证并行

	// 两组协程 其中一组负责写数据 一组负责读数据

	// 虽然有多个 goroutine 读数据之间的可以并发 写和读之间应该串行 读和读之间不能并行

	// 读写锁
	var num int
	var rwlock sync.RWMutex
	var wg sync.WaitGroup
	// write

	wg.Add(2)
	go func() {
		defer wg.Done()
		rwlock.Lock() // 加写锁 防止别的写锁 或者 读锁获取
		defer rwlock.Unlock()
		num = 12
		time.Sleep(time.Second * 5)
		fmt.Println("get write lock")
	}()
	time.Sleep(time.Second)
	// read
	go func() {
		defer wg.Done()

		for {
			rwlock.RLock() // 读锁 不会阻止别人的读
			fmt.Println(num)
			time.Sleep(500 * time.Millisecond)
			fmt.Println("get read lock")
			rwlock.RUnlock()
		}

	}()

	wg.Wait()
}
