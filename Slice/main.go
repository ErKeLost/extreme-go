package main

import "fmt"

func main() {

	// 切片就是数组 但是可以不定义长度

	var course []string
	fmt.Printf("%T\r\n", course)

	course = append(course, "go", "app")
	fmt.Println(course)

	// 切片的三种初始化方法  1.从数组中直接创建 2.使用 string{} 3.make

	allCourses := [4]string{"go", "grpc", "gin", "mysql"}

	courseSlice := allCourses[0:1]                // 左闭右开
	courseSlice2 := allCourses[0:len(allCourses)] // 左闭右开

	fmt.Println(courseSlice)
	fmt.Println(courseSlice2)

	// make 函数

	allCourses3 := make([]string, 3)
	allCourses3[0] = "c"

	fmt.Println(allCourses3)

	// 访问切片的元素 访问单个 和 访问多个
	fmt.Println(allCourses3[0])

	// [start:end] 如果只有start 没有 end 就是 从 开始到结尾的所有数据
	// 没有start 只有 end 从 0 到结尾
	// 没有 start 也没有 end  取所有
	fmt.Println(course[0:])
	fmt.Println(course[:])

	// 切片 里面数据的添加 append ...

	lesson := []string{"name", "age", "job"}
	lesson2 := []string{"ww", "ee"}
	lesson = append(lesson, "dodood")
	lesson = append(lesson, lesson2...)
	// 合并两个slice
	fmt.Println(lesson)

	// 如何删除 slice 中的元素

	//就是取元素
	mySlice := append(lesson[:2], lesson[3:]...)
	fmt.Println(mySlice)

	// 复制slice

	copySlice := lesson[:]

	fmt.Println(copySlice)

	sourcea := make([]string, len(lesson))
	copy(sourcea, lesson)

	fmt.Println(sourcea)
}
