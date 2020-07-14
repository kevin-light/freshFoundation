package main

import (
	"fmt"
	bmp "t01/packageT1/packageDemo/calc"  // 包的别名
)	// 合并导入

//import "fmt"
////	当你的代码都放在$GOPATH目录下的话，我们要从￥GOPATH/src后面路径开始写起
//import "t01/packageT1/packageDemo/calc"
// go语言不允许导入包不使用
func main()  {
	fmt.Println("hello")
	//ret := calc.Add(10,20)
	ret := bmp.Add(10,20)
	fmt.Println(ret)
}

func init()  {
	fmt.Println("main.init()")	//	init多用于初始化， main.init > calc.init
}