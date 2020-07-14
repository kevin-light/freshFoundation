package main

import (
	"fmt"
	"html/template"
	"net/http"
)
// 遇事不决 写注释  //  go web项目不要直接 run go

func sayHello(w http.ResponseWriter,r*http.Request)  {
	// 解析模板
	t,err:=template.ParseFiles("./hello.html")
	if err!= nil{
		fmt.Println("parse template failed err:",err)
		return
	}
	// 渲染模班
	name := "小明"
	err = t.Execute(w,name)
	if err!=nil{
		fmt.Println("render template failed err: ",err)
		return
	}

}
// 遇事不决 写注释
//func sayhello()  {
//
//}
func main()  {
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil{
		fmt.Println("http server start failed,err:",err)
		return
	}
}
