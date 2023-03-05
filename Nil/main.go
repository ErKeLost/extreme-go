package main

import "fmt"

func main() {
	/*
		不同类型的数据 零值 不一样
		bool false
		number 0
		string ""
		pointer nil
		slice nil
		map nil
		channel interface function nil

		struct 默认值不是nil 是具体字段的默认值
	*/
	type Person struct{}
	var ps []Person             // nil slice
	var ps2 = make([]Person, 0) // 这是个空的 slice slice 本质是一个 struct struct 空的是各个属性的值 不是 nil
	var ps3 = make([]Person, 0)

	//所以 make 也分 nil map  和 empty map
	// map 取值不会报错  赋值就会报错 make 方法 赋值不会报错
	if ps3 == nil {
		fmt.Println("nil slice ps3")
	}
	if ps2 == nil {
		fmt.Println("nil slice")
	}

	if ps == nil {
		fmt.Println(ps, "nil")
	}

}
