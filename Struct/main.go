package main

import (
	"fmt"
	"strconv"
)

type Per struct {
	name string
	ages int
}

// 方法接收器有两种形态 还有一种指针形态 如果你想要改变p的value
// 数据较大的时候也会使用指针的方式
func (p *Per) Print() {
	p.name = "9989456"
	fmt.Printf("name:%s, age: %d", p.name, p.ages)
}

//func (p Per) Print() {
//	fmt.Printf("name:%s, age: %d", p.name, p.ages)
//}

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

	// day 2
	fmt.Println("--------------")

	var p Person

	p.age = 20

	fmt.Println(p.name)
	fmt.Println(p.age)

	// 匿名结构体 匿名函数
	addres := struct {
		country string
	}{
		country: "通州",
	}
	fmt.Println(addres.country)

	// 结构体嵌套

	// 第一种嵌套方式
	type Stu struct {
		//p Per

		// 第二种嵌套方式 匿名嵌套
		Per // 就不需要在 s.p.name了

		score float32
	}

	s := Stu{
		Per{
			"erkelost", 999,
		},
		88.88,
	}
	fmt.Println(s)

	// 结构体绑定方法

	wwe := Per{
		"adny", 99,
	}
	wwe.Print()
}

type My int // 扩展其他方法 int 扩展一个 string 的方法

func (my My) string() string {
	return strconv.Itoa(int(my))
}
