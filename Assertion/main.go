package main

func add(a, b int) int {
	return a + b
}

func addInt32(a, b int32) int32 {
	return a + b
}

// 如果有多种其他类型 那么就要写好多函数
// 断言 转换 类型
func addT(a, b interface{}) interface{} {
	a1, b1 := a.(int), b.(int)
	return a1 + b1

	// 但是这样并不通用 如果用了 float32 还是会报错
}

// 通过 switch 进行类型判断
func reduce(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi

	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi

	case float64:
		ai, _ := a.(float64)
		bi, _ := b.(float64)
		return ai + bi
	case string:
		ai, _ := a.(string)
		bi, _ := b.(string)
		return ai + bi
	default:
		panic("not supported type")
	}
}
func main() {
	// 断言
	a, b := 1, 5
	add(a, b)
}
