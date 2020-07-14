package main

import "fmt"

// type NewInt int // 类型定义

// type MyInt = int //类型别名

// func main() {
// 	var a NewInt
// 	fmt.Printf("type:%T value:%v\n", a, a)
// 	var b MyInt
// 	fmt.Printf("type:%T value:%v\n", b, b)
// }

// //定义结构体
// type person struct {
// 	name, city string
// 	age        int
// 	// city string
// }
// func main() {
// 	var p person
// 	p.name = "小米"
// 	p.city = "上海"
// 	p.age = 19
// 	fmt.Println(p.age)
// 	fmt.Printf("%#v\n",p)
// }

// //匿名结构体
// func main() {
// 	var user struct {
// 		name string
// 		age  int
// 	}
// 	user.name = "小明"
// 	user.age = 18
// 	fmt.Printf("%#%\n", user)

// }

type person struct {
	name, city string
	age        int
}

// // // 结构体指针类型
func main() {

	// 	var p person
	// 	p.name = "小米"
	// 	p.city = "上海"
	// 	p.age = 19
	// 	fmt.Println(p.age)
	// 	fmt.Printf("p1=%#v\n",p)

	// // 使用new关键字对结构体进行实例化，得到的是结构体的地址

	// 	var p2 = new(person)
	// 	fmt.Printf("p2=%#v\n",p2)
	// 	// (*p2).name = "小明"
	// 	p2.name = "xiaoming"
	// 	p2.city = "上海"
	// 	p2.age = 17
	// 	fmt.Printf("p1=%#v\n",p2)

	// // 取结构体的地址实例化
	// 	p3 := &person{}
	// 	fmt.Printf("%T\n", p3)     //*main.person
	// 	fmt.Printf("p2=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	// 	p3.name = "七米"   //(*p3).name = "七米"
	// 	p3.age = 30
	// 	p3.city = "成都"
	// 	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}

	// // 1.键值对初始化
	// p4 := person{
	// 	name: "小明",
	// 	age:  18,
	// }
	// fmt.Printf("%#v\n", p4)

	// p5 := &person{
	// 	name: "浦",
	// 	city: "上海",
	// 	age:  123,
	// }
	// fmt.Printf("%v\n",p5)

	// 2值的列表初始化：必须初始化所有字段，顺序不能变
	// p6 := person{
	// "小JPG",
	// "shangh",
	// 15,
	// }
	// fmt.Printf("%#v\n",p6)

	//结构体构造函数，构造一个结构体实例的函数
	p1 := newPerson("小明", "上海", int(18))
	fmt.Printf("type:%T value:%#v\n", p1, p1)
}
// // 直接传递等于copy占用内存
// func newPerson(name, city string, age int) person {
// 	return person{
// 		name: name,
// 		city: city,
// 		age:  age,
// 	}
// }
// 传递指针节省内存（常用）
func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}