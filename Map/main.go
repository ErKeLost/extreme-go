package main

import (
	"container/list"
	"fmt"
)

func main() {
	// map 是一个 key value 的无序集合 主要是用来查询 不想 数组 和 slice 需要遍历查询

	var mapA = map[string]string{
		"go":   "go 全栈开发",
		"grpc": "grpc 冲冲冲",
	}
	fmt.Println(mapA)
	// map 类型想要设置值 必须要先初始化

	mapB := map[int]string{}
	mapB[0] = "我是 mysql"
	fmt.Println(mapB)

	//make 方法 内置函数 初始化 slice map channel
	var mapSource = make(map[string]string, 66)

	fmt.Println(mapSource)

	// map 必须初始化才能使用 map[string]int{}  make(map[string]string, 3)
	// 但是slice 可以不用初始化 就可以赋值

	// map的遍历

	for key, value := range mapA {
		fmt.Println(value, key)
	}

	for value := range mapA {
		fmt.Println(value)
	}

	// 无序集合 每次循环打印出来的都是不一样的 顺序

	// map 获取元素和删除元素 判断元素是否存在

	fmt.Println(mapA["go"])

	d, ok := mapA["java"]

	if !(ok) {
		fmt.Println("not in mapA")
	} else {
		fmt.Println(d)
	}

	delete(mapA, "go")
	fmt.Println(mapA)

	// map不是线程安全的  多线程 操作一个 map 会报错

	// list 链表

	//slice在添加数据 会不停的 扩容 拷贝

	// list 的基本用法

	var myList list.List

	myList.PushBack("go")
	myList.PushBack("grpc")
	myList.PushBack("mysql")
	fmt.Println(myList.Front())

	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	// 正向遍历 和 反向遍历

	//myList := list.New()
}
