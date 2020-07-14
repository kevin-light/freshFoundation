package main

import "fmt"

// 变量
var x = 10
var n = "小王子"

// 常量
const (
	n1 = 10
	n2 = 3.14
)

// 变量
func main() {
	// 标准声明方式
	var name string
	var age int
	var isok bool
	fmt.Println(name, age, isok)
	// 批量声明
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	// 声明变量同时指定初始值
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Println(name1, age1)
	// 类型推到(常用)，让编译器根据变量的初始值推到类型
	var name3 = "小王子3"
	var age3 = 18
	fmt.Println(name3, age3)
	// 短变量声明
	m := 10
	fmt.Println(m)

}

func foo() (int, string) {
	return 10, "小王子"
}

func main1() {
	// var x = 0
	x, _= foo()
}
