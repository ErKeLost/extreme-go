package main

import (
	//u "extreme-go/Package-a/User" // 别名
	. "extreme-go/Package-a/User" // 全局导入 不需要.
	_ "extreme-go/Package-a/User" // 导入必须使用 这样就可以不使用 就进行初始化代码
	"fmt"
)

func main() {
	c := User{
		Name: "what",
	}

	fmt.Println(c)
}
