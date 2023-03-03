package main

import "fmt"

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
}
