package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 转义符
	//courseName := "我是\"奥特曼\""
	courseName := `我是"奥特曼"`
	courseName2 := "我是\r\n\"奥特曼\""
	fmt.Println(courseName)
	fmt.Println(courseName2)

	/*
		换行符‘\n’和回车符‘\r’

		顾名思义，换行符就是另起一行，回车符就是回到一行的开头，所以我们平时编写文件的回车符应该确切来说叫做回车换行符

		'\n' 10 换行（newline）
		'\r' 13 回车（return）
	*/

	// 格式化输出
	name := "erkelost"
	out := "username\r\n" + name
	fmt.Println(out)

	// %v 打印格式
	/*
		%#v  go 语法打印 打印格式语法
		%T 类型打印
		%d 十进制打印
		%+d 正负号
		%4d pad空格
		%-4d
		%b 二进制
	*/

	// 通过 string builder 字符串拼接 高性能
	const (
		name1   = "erkelost"
		age1    = 99
		address = "us"
		phone   = "13854214521"
	)
	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(name1)
	builder.WriteString("age")
	builder.WriteString(strconv.Itoa(age1))
	builder.WriteString("address")
	builder.WriteString(address)

	fmt.Println(builder.String())

	// 字符串其他常用功能

	//字符串的比较
	a1 := "hello"
	b1 := "hello"

	fmt.Println(a1 != b1)

	boola := strings.Contains(a1, "hell")
	fmt.Println(boola)

	// 出现次数
	fmt.Println(strings.Count(a1, "o"))
	fmt.Println(strings.Split(a1, "l"))
	fmt.Println(strings.HasPrefix(a1, "l"))
	fmt.Println(strings.HasSuffix(a1, "l"))
	fmt.Println(strings.Index(a1, "l"))

	fmt.Println(strings.ToUpper(a1))
	fmt.Println(strings.Trim(a1, "hel"))
}
