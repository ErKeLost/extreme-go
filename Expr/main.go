package main

import "fmt"

func main() {
	// 最近本的运算符  加减乘除
	var a, b = 1, 2
	fmt.Println(a + b)

	var astr, bstr = "hello", "world"
	fmt.Println(astr + bstr)

	a++
	a += 1
	a = a + 1

	// 逻辑运算符
	var abool, bbool = true, false
	if abool && bbool {
		fmt.Println()
	} else {
		fmt.Println(666)
	}

	// 位运算符
	// p & q 与运算符
	// p | q 或运算符
	// p ^ q 异或 运算符  值相同 就是 0 值不同就是1

	var A, B = 60, 54
	fmt.Println(A & B)
}
