package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 2. 新的需求 主动退出监控程序 可以使用共享变量

var stop bool

// 使用channel

var stop1 = make(chan struct{})

func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		//if stop {
		//	break
		//}
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu 监控")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("cpu 信息打印")
		}

	}

}

// 父给字 传递信息中断子goroutine 使用context

func main() {
	// 渐进式 为什么要使用 context

	// 有一个 goroutine 监听cpu 信息

	// 防止 主goroutine 退出 使用 waitGroup
	wg.Add(1)

	// context 包提供三种函数 withCancel withtimeout withvalue
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx1)
	go cpuInfo(ctx2)
	time.Sleep(6 * time.Second)
	cancel1() // 父的cancel 方法 会递归调用 🌲 的所有子cancel方法
	wg.Wait()
	fmt.Println("监控完成")
}
