package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	gender string	// 首字母小写 模板不能访问 !!!!!!
	Age int
}
func sayhello(w http.ResponseWriter, r *http.Request)  {
	//定义模班
	//解析模板
	t,err:=template.ParseFiles("./hello.html")
	if err!=nil{
		fmt.Println("parse template failed err:",err)
		return
	}
	//渲染模板
	u1 := User{
		Name: "小明",
		gender: "男",
		Age: 12,
	}
	m1 := map[string]interface{}{
		"Name":"小红",
		"gender":"女",
		"age":17,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w,map[string]interface{}{  	// 利用给定数据渲染模板，并将结果写入w

		"u1":u1,
		"m1":m1,
		"hobbyList":hobbyList,
	})

}

func main()  {
	http.HandleFunc("/",sayhello)
	err := http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("http server start failed,err:",err)
		return
	}
}