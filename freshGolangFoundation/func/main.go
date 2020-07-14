package main

import "fmt"

func intSum(a int, b int) int {
	ret := a + b
	return ret
}

func intSum2(a int, b int) (ret int) {
	ret = a + b
	return ret
}

//  ... 函数接收可变参数
func intSum3(a ...int) int {
	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	return ret
}

func intSum5(a int, b ...int) int {
	ret := a
	for _, v := range b {
		ret = ret + v
	}
	return ret
}

// //定义多个返回值(sum,sub int)
// func calc(a, b int) (sum, sub int) {
// 	sum = a + b
// 	sub = a - b
// 	return
// }

func main() {
	// r := intSum3(10, 20)
	// fmt.Println(r)

	// r:= intSum3(10)
	// r1:= intSum3(10,20)
	// r2:= intSum3(10,20,30,70)

	// x, y := calc(10, 20)
	// fmt.Println(x, y)

	// testGlobal()
	// abc := testGlobal //函数传递
	// abc()

	// r1 := calc(10, 20, add)
	r2 := calc(10, 20, sub) //函数作为参数传递
	fmt.Println(r2)
}

//定义全局变量num和局部变量优先级，函数内部局部变量优先级高于全局变量
//外部变量访问局部变量： 通过返回值访问
// for 循环的语句只在循环内生效
var num = 10

func testGlobal() {
	num := 100
	fmt.Println("变量num", num)
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func calc(x, y int, op func(int, int) int) int {  
	return op(x, y)

}
