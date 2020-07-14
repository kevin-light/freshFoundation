package calc

import "fmt"

var name = "小王子"  // 首字母小写外部包不能访问

func add(x,y int) int {	 // 首字母小写外部包不能访问
	return x+y
}


var Name = "小王子"  // 首字母小写外部包不能访问

func Add(x,y int) int {	 // 首字母小写外部包不能访问
	return x+y
}

// init函数在导入的时候自动执行，init函数没有参数也没有返回值
func init(){
	fmt.Println("calc.init()")
}