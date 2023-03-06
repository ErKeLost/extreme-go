package main

import "fmt"

func main() {
	// 接口 鸭子类型
	// go 语言中到处都是鸭子类型

	/*
		当看到一只鸟走起来像鸭子
		游泳像鸭子
		叫起来像鸭子
		那么这只鸟就是鸭子
	*/

	//duck :=
	//javaScript 举例
	//class Person {
	//	constructor(public name: string)
	//}
	//
	//class Dog {
	//	constructor(public name: string)
	//}
	//
	//function print(p: Person) {}
	//
	//print(new Dog('erkelost'))

	var d Duck = &psdDuck{}
	d.Walk()

	// 多接口的实现 一个 结构体实现两个接口
	//var mm MyWriter = &Wc{}
	//var ww MyClose = &Wc{}

	// 接口的嵌套

	var wwe MyWc = &oow{}

	wwe.Close()
}

type MyWriter interface {
	Write(string) error
}

type MyClose interface {
	Close() error
}

type MyWc interface {
	MyClose
	MyWriter
}

type oow struct {
}

func (o oow) Close() error {
	//TODO implement me
	panic("implement Close")
}

func (o oow) Write(s string) error {
	//TODO implement me
	panic("implement Write")
}

type Wc struct {
}

func (wc *Wc) Write(string) error {
	fmt.Println("write string")
	return nil
}

func (wc *Wc) Close() error {
	fmt.Println("Close string")
	return nil
}

// 接口的定义 接口的定义没有用 接口往往都需要实现

type Duck interface {
	// 方法

	Gaga()
	Walk()
	Swimming()
}

type psdDuck struct {
	legs int
}

func (pd *psdDuck) Gaga() {
	fmt.Println("gagaga")
}

func (pd *psdDuck) Walk() {
	fmt.Println("walk ....")
}

func (pd *psdDuck) Swimming() {
	fmt.Println("Swimming ....")
}
