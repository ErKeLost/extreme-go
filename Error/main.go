package main

import (
	"errors"
	"fmt"
)

func main() {
	// 错误处理
	// 开发函数的人需要有一个返回值去告诉调用者是否成功
	if _, err := A(); err != nil {
		fmt.Println("成功了", err)
	}
}

func A() (int, error) {

	// 被动触发panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("主动拦截 panic")
		}
	}()

	//defer 要放在 panic 之前调用
	var names map[string]string
	names["go"] = "gogogogo"
	panic("fuck") // 一般服务启动过程中 日志文件存在 mysql 连通 配置文件没有问题 任何文件不满足 我们主动退出服务

	// recover 捕获到 panic 防止程序主动退出
	return 0, errors.New("this is error")
}

// panic 导致程序推出 exit code  go 语言不推荐随便使用 panic

// defer 需要放在 panic 之前定义 另外 recover 只有在defer 调用的函数中 生效
// recover 处理异常之后 逻辑不会恢复到 panic 那个点
// 多个defer 会形成zhan 然后 依次执行
