package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基本数据类型
	/*
		bool 类型
		值可以是常量 true 或者 false
		b := false
	*/

	/*
		int 8
	*/

	var a int8
	var b int16
	var c int32
	var d int64
	//var ua uint8
	//var ub uint16
	//var uc uint32
	//var ud uint64
	fmt.Println(a, b, c, d)

	var q float32 // 大约是 3.4e38
	var w float64 // 大约是 1.8e308

	fmt.Println(q, w)

	// int 类型 是动态类型 根据 计算机来的

	//byte类型
	var dd byte // uint8 别名 是 byte类型 主要是用来 存放字符的
	fmt.Println(dd)

	var cc byte = 'a'
	ccc := 'b'
	fmt.Println(cc)
	fmt.Printf("cc=%c", cc)
	fmt.Printf("cc=%c", ccc)

	fmt.Println()
	// rune 本质是一个 int32 他主要是用来 放字符的
	//var c2 rune
	//c2 = "我是一个小男孩"
	//fmt.Printf("c2=%c", c2)

	// 数据类型的转换

	//int 类型之间的转换

	//convert

	var a1 int8 = 12
	var b1 = uint8(a1)

	var c1 float32 = 3.14
	var d1 = int32(c1)

	fmt.Println(b1, d1)

	type IT = int // 类型 定义别名之后 不能进行类型转换
	//var ff IT = 999

	//字符串转数字
	var istr = "999"
	myint, err := strconv.Atoi(istr)

	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint, istr)

	myi := 32
	msyi := strconv.Itoa(myi)
	fmt.Println(msyi)
}
