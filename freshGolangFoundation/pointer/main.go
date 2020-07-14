package main

import "fmt"

// func main() {
// 	a := 10
// 	b := &a  // 取变量a的地址(指针)，将指针保存到b中
// 	fmt.Printf("a:%d ptr:%p\n", a, &a)
// 	fmt.Printf("b:%p type:%T\n", b, b)
// 	fmt.Println(&b)
// 	c:=*b 	//根据指针去内存取值
// 	fmt.Printf("type of c:%T\n",c)
// 	fmt.Printf("value of c:%v\n",c)
// }

func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 100

}
func main() {
	// a := 10
	// modify1(a)
	// fmt.Println(a)
	// modify2(&a)  	//传递指针，函数给指针赋值
	// fmt.Println(a)

	var a *int
	a = new(int)  	//初始化int，不常用，常用 a := 10
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int,10)
	b["one"] = 100
	fmt.Println(b)
}
