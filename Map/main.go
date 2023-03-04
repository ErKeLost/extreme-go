package main

import "fmt"

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
}
