package main

import (
	"fmt"
	"sync"
)

func add(a int, b int) (int, error) {
	return a + b, nil
}

func reduce(a, b int) (sum int, err error) {
	sum = a - b
	return sum, err
}

func some(item ...int) (sum int) {
	for _, value := range item {
		sum += value
	}
	return sum
}

func autoIncrement() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}

func main() {
	// go 函数支持普通函数 匿名函数 闭包
	// go 函数是一等公民 函数本身可以当作变量 函数可以满足接口
	res, _ := add(5, 666)

	fmt.Println(res)
	fmt.Println(some(1, 2, 3, 45, 45, 42, 12))
	// go 语言全部都是值传递

	// 函数一等公民的特性
	funVar := add
	fmt.Println(funVar(66, 88))

	// 闭包特性
	//需求
	/*
		希望函数每次返回值都是增加之后的值
	*/
	next := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(next())
	}
	// 链接数据库 打开文件 锁 做完操作要关系文件 解锁

	var mu sync.Mutex
	mu.Lock()
	// 不需要 try  finally 释放锁
	defer mu.Unlock()

	// defer 后面的代码 会放在函数return之前执行
	// 打开文件 释放文件 defer 先后顺序
	// defer 是一个栈的概念先进后出
	ret := deferReturn()
	fmt.Println(ret)
	defer func() {

	}()
}

// 函数中的 defer
func deferReturn() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}
