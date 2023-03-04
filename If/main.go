package main

import "fmt"

func main() {
	//if (1 == 1) && (2 == 22) {
	//	fmt.Println("you win")
	//} else {
	//	fmt.Println("you lose")
	//}

	// for 循环
	/*
		for  init ; condition ; post {}
	*/
	//for i := 0; i < 3; i++ {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println(i + 99)
	//}

	//var i int
	//for i < 3 {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println(i + 99)
	//	i++
	//}

	//for ; i < 10; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d*%d=%d ", i, j, i*j)
	//	}
	//	fmt.Println()
	//}
	//
	//// forRange
	//
	//for key, value := range "我是一个奥特曼" {
	//	fmt.Println(key, value)
	//	fmt.Printf("%d, %c\r\n", key, value)
	//}
	//
	//for _, value := range "我是一个奥特曼" {
	//	fmt.Println(value)
	//	fmt.Printf(" %c\r\n", value)
	//}
	//round := 1
	//for {
	//	time.Sleep(1 * time.Second)
	//
	//	round++
	//	if round == 6 {
	//		continue
	//	}
	//	fmt.Println("睡觉")
	//
	//	if round > 10 {
	//		break
	//	}
	//}

	// goto 语句 跳到指定代码中运行

	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			if j == 2 {
				goto over
			}

			fmt.Println(i, j)
		}
	}
over:
	fmt.Println(666)
	// 错误处理的考虑

	// Switch
	var score int
	switch score {
	case 99:
		fmt.Println(99999)
	default:
		fmt.Println(565656)
	}

}
