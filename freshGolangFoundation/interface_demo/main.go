package main

import (
	"fmt"
)

type dog struct {

}
type cat struct {

}
type person struct {
	name string
}

func (p person) say() {
	fmt.Println("aaaaaa!")
}
func (d dog) say() {
	fmt.Println("wangwangwang")
}

func (c cat) say()  {
	fmt.Println("miaomiaomiao")
}
//  接口不管是什么类型，只管它实现什么方法
// 定义各类型，一个抽象的类型
type sayer interface {
	say()
}
func da(arg sayer)  {
	arg.say()		// 传进来的，打他，触发say
}

type move interface {
	move()
}
type psn struct {
	name string
	age int
}

//func (p psn) move() {
//	fmt.Printf("%s再跑。。。",p.name)
//}
// 使用指针接收者实现接口：只有类型的指针可以保存到接口
func (p *psn) move() {
	fmt.Printf("%s再跑。。。",p.name)
}
func (p *psn) say()  {
	fmt.Printf("%s再叫",p.name)
}
//接口嵌套
type  animal interface {
	move
	sayer
}
func main()  {
	//c1 := cat{}
	//da(c1)
	//p1 := person{name: "小明"}
	//da(p1)
	//var s sayer
	//c2 := dog{}
	//s = c2
	//fmt.Println(s)

	//var m move
	//p1 := psn{
	//	name: "小明",
	//	age: 18,
	//}
	//m=p1
	//m.move()
	//fmt.Println(m)
	//
	p2 := &psn{		// p2是psn类型的指针
		name: "小红",
		age: 17,
	}
	//m=p2
	//m.move()
	//fmt.Println(m)

	var s sayer
	s = p2
	s.say()
	fmt.Println(s)

	var a animal
	a = p2
	a.move()
	a.say()
	fmt.Println(a)

	// 定义一个空接口，空接口变量可以存储任意类型
	var x interface{}
	s1 := "沙河"
	x = s1
	fmt.Printf("type:%T value:%v\n",x,x)
	i := 100
	x = i
	//fmt.Printf("type:%T value:%v\n",x,x)
	//var m1 = make(map[string]interface{},16)
	//m1["name"] = "小明"
	//m1["age"] = 18
	//m1["bobby"] = []string{"篮球","双色球"}
	//fmt.Println(m1)

	ret, ok := x.(string)	// 类型断言
	if !ok {
		fmt.Println("不是字符串类型")
	}else {
		fmt.Println("是字符串类型", ret)
	}

	// 使用switch进行断言
	switch v := x.(type) {
	case string:
		fmt.Println("x is sting",v)
	case int:
		fmt.Printf("x is a int %v \n",v)
	default:
		fmt.Println("unsupport type")

	}

}