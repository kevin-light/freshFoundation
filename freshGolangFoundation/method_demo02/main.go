package main

import "fmt"


// 结构体继承
type animal struct{
	name string
}
func (a *animal) move(){
	fmt.Printf("%s惠东",a.name)
}
type dog struct{
	feet int64
	*animal	//	匿名嵌套结构体的指针
}
func (d *dog) wang(){
	fmt.Printf("%s会汪汪、\n",d.name)
}
func main(){
	d1 := &dog{
		feet:3,
		animal:&animal{
			name:"lele",
		},
	}
	d1.wang()
	d1.move()
}

// // 嵌套结构体的字段冲突
// type address struct {
// 	provice    string
// 	city       string
// 	updateTime string
// }
// type Emai struct {
// 	addr       string
// 	updateTime string
// }
// type person struct {
// 	name   string
// 	gender string
// 	age    int
// 	address
// 	Emai
// }

// func main() {
// 	p1 := person{
// 		name:   "小红",
// 		gender: "半",
// 		age:    17,
// 		address: address{
// 			provice:    "山东",
// 			city:       "上海",
// 			updateTime: "1001",
// 		},
// 		Emai: Emai{
// 			addr:       "www.baidu,xon",
// 			updateTime: "303030",
// 		},
// 	}
// 	fmt.Println(p1.address.updateTime)
// 	fmt.Println(p1.Emai.updateTime)
// }

// // Person 嵌套结构体
// type address struct {
// 	provice string
// 	city string
// }
// type Person struct {
// 	name,gender string
// 	age int
// 	// address address
// 	address  // 嵌套另一个结构体
// }
// func main(){
// 	p1 := Person{
// 		name:"小明",
// 		gender:"男",
// 		age:16,
// 		address:address{
// 			provice:"上海",
// 			city:"chang",
// 		},
// 	}
// 	fmt.Printf("%#v\n",p1)
// 	fmt.Println(p1.name,p1.gender,p1.age,p1.address)
// 	fmt.Println(p1.address.provice)
// 	fmt.Println(p1.provice)
// }

// //  Person 结构体的匿名字段，取值有限制，局现象
// type Person struct {
// 	string
// 	int
// }

// func main() {
// 	p1 := Person{
// 		"小敏",
// 		15,
// 	}
// 	fmt.Println(p1)
// 	fmt.Println(p1.string, p1.int)
// }
