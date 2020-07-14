package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性与json序列化
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
type student struct {
	ID int
	Name string
}

// student 构造函数
func newStudent(id int, name string) student{
	return student{
		ID: id,
		Name: name,
	}
}

type class struct {
	//Title string		//首字母小写其他包不能访问
	Title string 	`json:"title"`		//通过反射的机制读取首字母小写的title
	Student []student  // 构造切片类型
}
func main()  {
	//创建一个班级变量
	c1 := class{
		Title: "高一级",
		Student: make([]student, 0, 20),		// make对切片初始化
	}
	//网班级添加学生
	for i:=0;i<10;i++{
		//构造10个学生
		tempStu := newStudent(1,fmt.Sprintf("str%02d",i))
		c1.Student=append(c1.Student,tempStu)
	}
	fmt.Printf("%#v\n", c1)
	// json序列化：go语言的数据 -》 json格式的字符串
	data,err:=json.Marshal(c1)
	if err != nil{
		fmt.Println("json marshal failed err:",err)
		return
	}
	fmt.Printf("%T\n",data)
	// json反序列化： json字符串 -》 go语言中的数据
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c2 := &class{}
	err = json.Unmarshal([]byte(str),&c2)
	if err != nil{
		fmt.Println("json unmarshal err:",err)
		return
	}
	fmt.Printf("%#v\n",c1)
}