package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 切片就是数组 但是可以不定义长度

	//var course []string
	//fmt.Printf("%T\r\n", course)
	//
	//course = append(course, "go", "app")
	//fmt.Println(course)
	//
	//// 切片的三种初始化方法  1.从数组中直接创建 2.使用 string{} 3.make
	//
	//allCourses := [4]string{"go", "grpc", "gin", "mysql"}
	//
	//courseSlice := allCourses[0:1]                // 左闭右开
	//courseSlice2 := allCourses[0:len(allCourses)] // 左闭右开
	//
	//fmt.Println(courseSlice)
	//fmt.Println(courseSlice2)
	//
	//// make 函数
	//
	//allCourses3 := make([]string, 3)
	//allCourses3[0] = "c"
	//
	//fmt.Println(allCourses3)
	//
	//// 访问切片的元素 访问单个 和 访问多个
	//fmt.Println(allCourses3[0])
	//
	//// [start:end] 如果只有start 没有 end 就是 从 开始到结尾的所有数据
	//// 没有start 只有 end 从 0 到结尾
	//// 没有 start 也没有 end  取所有
	//fmt.Println(course[0:])
	//fmt.Println(course[:])
	//
	//// 切片 里面数据的添加 append ...
	//
	//lesson := []string{"name", "age", "job"}
	//lesson2 := []string{"ww", "ee"}
	//lesson = append(lesson, "dodood")
	//lesson = append(lesson, lesson2...)
	//// 合并两个slice
	//fmt.Println(lesson)
	//
	//// 如何删除 slice 中的元素
	//
	////就是取元素
	//mySlice := append(lesson[:2], lesson[3:]...)
	//fmt.Println(mySlice)
	//
	//// 复制slice
	//
	//copySlice := lesson[:]
	//
	//fmt.Println(copySlice)
	//
	//sourcea := make([]string, len(lesson))
	//copy(sourcea, lesson)
	//
	//fmt.Println(sourcea)
	// 切片的原理

	// go slice 在函数参数传递的时候是值传递还是引用传递  值传递
	// 效果呈现除了引用传递的效果
	courses := []string{"go", "grpc", "gin"}
	printSplice(courses)

	fmt.Println(courses)
	// courses 的值发生了变化 效果好像是因为引用传递

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := data[1:6]
	s2 := data[2:7]
	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	s2[0] = 22

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s1)
	fmt.Println(s2)
}

func printSplice(data []string) {
	data[0] = "java"

	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}

	// 但是这个append 又接受不到 不能给 data fuzhi
}

// slice 底层原理  slice 本质是一个 struct 结构体

// 正常修改值 引用传递 但是 如果 你append 他就会扩容成倍扩容 需要 复制一份值 就是值传递

//就是 如果你 append 他会扩容 重新生成一份地址
