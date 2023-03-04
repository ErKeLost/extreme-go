package main

import (
	"fmt"
	"strconv"
)

func main() {
	// type 关键字

	/*
		type 有几种用途
		1. 定义结构体
		2. 定义接口
		3. 定义类型别名
		4. 自定义类型
	*/

	// 别名的是为了更好的理解代码

	type MyInt = int8
	type MyInt2 = int

	var a MyInt2 = 66
	fmt.Printf("%T\r\n", a)

	var b int = 12

	fmt.Println(a + b)

	// 类型定义 自定义类型

	// 自定义类型 基于已有的类型
	var i My = 666

	fmt.Printf("%T\r\n", i)

	// 想放 多个persion的信息放到一个集合中
	var persons [][]string

	persons = append(persons, []string{"erkelost", "18", "muke", "1.88"})
	// 这样如果我们需要定义不同类型 我们就可以使用 结构体 目前就是string 类型的
	fmt.Println(persons)

	type Person struct {
		name    string
		age     int
		address string
		height  float32
	}

	// 如何初始化结构体
	p1 := Person{"adny", 88, "github", 5.22}
	p2 := Person{
		name:    "erkelost",
		age:     88,
		address: "gitlab",
		height:  2.222,
	}

	fmt.Println("p1", p1)
	fmt.Println("p2", p2)

	var persion []Person

	persion = append(persion, p1)
	persion = append(persion, p2)

	fmt.Println(persion)

	ww := []Person{
		{name: "nuoke",
			age:     99,
			address: "454564",
			height:  5.1212},
	}
	fmt.Println(ww)
}

type My int // 扩展其他方法 int 扩展一个 string 的方法

func (my My) string() string {
	return strconv.Itoa(int(my))
}
