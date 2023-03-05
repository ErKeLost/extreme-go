package main

import "fmt"

type Page struct {
	name string
}

func changeName(p *Page) {
	p.name = "1231231"
}

func main() {
	// 什么叫指针  我希望结构体传值的时候 我在函数中修改的值可以反应到变量中
	ww := Page{
		name: "erkelost",
	}

	changeName(&ww)

	fmt.Println(ww)

	po := &Page{
		name: "ppp",
	}
	// go 语言限制了指针的运算
	po.name = "123123123"

	// 指针的初始化
	var a int = 100

	//b := &a

	// 指针需要初始化
	var b *int
	fmt.Println(b, a)

	// 指针需要初始化

	ps := &Page{} // 第一种初始化指针

	var empty Page

	pi := &empty // 第二种初始化指针

	fmt.Println(ps, pi)

	var pp *Page = new(Page) // 第三种初始化指针

	fmt.Println(pp)

	// 初始化关键字 map channel slice 初始化推荐使用 make 方法
	// map 必须初始化 指针初始化推荐new 函数 指针必须要初始化

	// 通过指针交换两个值
	aw, bw := 1, 2

	swap(&aw, &bw)
	fmt.Println(aw, bw)
}

func swap(a, b *int) {
	//a, b = b, a

	// 这样是改不了的

	// 创建一个临时值
	t := *a
	*a = *b
	*b = t

	fmt.Println(a, 456)
	fmt.Println(*a, 789)
}
