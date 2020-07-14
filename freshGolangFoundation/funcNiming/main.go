package main

import "fmt"

func main() {
	// 闭包 = 函数+外层变量的引用
	// func() {
	// 	fmt.Println("匿名函数1")
	// }() // 立即执行func函数

	// r := a("hello")
	// r()
	x, y := calc(100)
	ret1 := x(200)
	fmt.Println(ret1)
	ret2 := y(200)
	fmt.Println(ret2)

	funcA()
	funcB()
	funcC()

}

//定义一个函数返回一个函数
func a(name string) func() {

	return func() {
		fmt.Println("匿名函数2", name)
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// recover()必须搭配defer使用。defer一定要在可能引发panic的语句之前定义
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

// 内置函数		介绍
// close	主要用来关闭channel
// len		用来求长度，比如string、array、slice、map、channel
// new		用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
// make		用来分配内存，主要用来分配引用类型，比如chan、map、slice
// append	用来追加元素到数组、slice中
// panic和recover	用来做错误处理
