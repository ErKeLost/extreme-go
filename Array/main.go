package main

import "fmt"

func main() {
	// go 中的集合数据结构 数组 切片 map list

	// 数组的定义

	//var Names [3]int = [1,2,3]
	var a [3]string
	a[0] = "go"
	a[1] = "grpc"
	a[2] = "crpc"
	fmt.Println(a)
	fmt.Printf("%T", a)
	fmt.Println()
	// []string && [3]string 这是两种不同的类型
	for _, value := range a {
		fmt.Println(value)
	}

	var b [4]string = [4]string{"a", "vb", "vvb", "we"}
	c := [4]string{"a", "vb", "vvb", "we"}
	d := [4]string{3: "we"}
	fmt.Println(c)
	fmt.Println(d)
	for _, value := range b {
		fmt.Println(value)
	}

	// 不知道定义多少长度

	dd := [...]string{"go", "rust", "javascript"}
	fmt.Println(dd)

	// 多维数组

	ee := [3][4]string{}
	ee[0] = [4]string{"go", "grpc", "web", "java"}
	fmt.Println(ee)
}
