package main

import "fmt"

// 全局变量 局部变量
var name = "erkelost"
var age = 999
var ok = false

//var (
//	name = "name"
//	age  = 666
//	ok   = true
//)

func main() {
	fmt.Println("---------变量")
	// go 是静态语言 静态语言和动态语言变量差异很大
	//编译必须先定义后使用 必须有类型 定下来之后不能改变
	//var name int = "erkelost"
	name := "adny"
	fmt.Println(name)

	// 多变量定义

	var user1, user2, user3 = "a1", false, "a3"
	fmt.Println(user3, user2, user1)

	// 变量必须先定义才能使用
	// 全局变量优先级高
	// 变量名尽量不要冲突  简洁变量定义不能用于全局
	//obj := 999
	//变量有0值 就是默认值
	var age2 int
	fmt.Println(age2)
	// 变量定义了就要使用 bool 类型的值 默认是false

	// 常量
	fmt.Println("---------常量")
	// 常量就是 定义的时候就是指定的值 不能修改
	const PI float32 = 3.1415926 // 显示定义
	const PII = 3.1415926857     // 隐示定义

	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)

	// 分组定义的时候规则

	/*
		常量类型只可以定义 bool 数值 字符串 不曾使用的常量 没有强制使用的要求
		显示指定类型的时候 必须确保左右类型的值一致
	*/

	const (
		x int = 16
		z
		s = "adny"
		y
	)
	// 没有定义 会延用 上面定义变量的值
	const a = 666
	fmt.Println(x, y, s, z, a)

	fmt.Println("--------------IOTA")

	// 特殊常量  iota 可以认为是一个被编译器修改的常量
	//const (
	//	ERR1 = iota
	//	ERR2 = iota
	//	ERR3 = iota
	//)

	const (
		ERR1 = iota + 2
		ERR2
		ERR3
		ERR4 = "adny"
		ERR5
		ERR6 = iota
	)

	// 只打印 iota 计数器 每次加1 不管你其他是什么值 后续会自动递增
	// 自增类型 默认是 int 类型 可以简化 const 类型的定义
	// 每次出现 const 关键字的时候 iota 归零

	fmt.Println(ERR1, ERR2, ERR3)

	fmt.Println("-----------匿名变量")

	// 匿名变量 定义变量的时候 go 必须使用

	//var name int ERROR
	var _ int
	// 因为变量必须要使用 所以我们可以打印一下 但是这样 有点刻意 所以我们可以使用 匿名变量
	//r1, ok := a()
	//_, ok := a(22, false)
	//if ok {
	//	fmt.Println(r1, ok)
	//}

	// 变量的作用域

	fmt.Println("----------变量的作用域")

	var mainName = "main"

	{
		localName := "user"
		fmt.Println(localName, mainName)
	}
}

//func a(int, bool) {
//	return false
//}
