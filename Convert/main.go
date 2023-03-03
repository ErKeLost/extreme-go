package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int8 = 12
	var b uint8 = uint8(a)

	var f float32 = 3.14
	var c = int32(f)

	fmt.Println(b, c)

	var f64 = float64(a)

	fmt.Println(f64)

	type IT int

	//var e IT = a

	// 类型要求很严格 type 定义别名之后 就不能转换

	//字符串转数字
	var isStr = "12我是"

	ww, err := strconv.Atoi(isStr)

	if err != nil {
		fmt.Println("convert error")
	}

	fmt.Println(ww)

	// 字符串转float32 bool
	nn, err := strconv.ParseFloat("3.1415926", 64)
	fmt.Println(nn)

	// 字符串转其他类型都有可能会出错的 因为你不知道他会转换成什么类型

	oo, err := strconv.ParseInt("12", 8, 64)
	fmt.Println(oo)

	ee, err := strconv.ParseBool("false")
	fmt.Println(ee)

	// 基本类型转string

	boolString := strconv.FormatBool(true)
	fmt.Println(boolString)

	floatStr := strconv.FormatFloat(3.1415926, 'E', -1, 64)
	fmt.Println(floatStr)

	intStr := strconv.FormatInt(42, 16)
	fmt.Println(intStr)

}
