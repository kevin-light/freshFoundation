package main

import "fmt"

// 方法的定义实例

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 是一个Person类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 是为Person定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是babab\n", p.name)
}

// SetAge 是一个指针 修改年龄的方法；指针接收者是指接收者的类型是指针
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}
// SetAge2 是一个值 修改年龄的方法；值 接收者是指接收者的类型是值
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}
// 需要修改接收者中的值
// 接收者是拷贝代价比较大的大对象
// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。



// 基于内置的类型造一个我们自己的类型
type myInt int
func (i myInt) sayHi(){
	fmt.Println("hi")
}

// 结构体的匿名字段

func main() {
	p1 := NewPerson("小红",int8(13))
	// p1.Dream()
	// // (*p1).Dream()

	// p1.SetAge(int8(15))
	// fmt.Println(p1.age)

	p1.SetAge2(int8(16))  //函数值copy
	fmt.Println(p1.age)
	
	var  i0 myInt
	i0 = 100
	fmt.Println(i0)

}
