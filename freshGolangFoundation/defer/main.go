package main

import "fmt"

// defer定义的函数逆序执行，先被defer的语句之后被执行，后被defer的语句后被执行
// defer的延迟特性，方便处理资源释放：比如：资源清理、文件关闭、解锁及记录时间,
// func main() {
// 	fmt.Println("start")
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3)
// 	fmt.Println("end")
// }

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
