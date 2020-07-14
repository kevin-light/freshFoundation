package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter,r *http.Request)  {
	// 定义一个函数框,
	kua := func(name string)(string,error){
		return name + "年轻帅气",nil		// 只能返回一个参数，或者第二参数必须是error类型
	}
	//定义模板
	t := template.New("index.html")	// 创建一个名字是f的模板对象, 名字一定要和ParseFiles内的函数对应上
	// 告诉模板引擎，传入一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua99":kua,
	})
	//解析模板
	_,err := t.ParseFiles("./index.html")
	if err != nil{
		fmt.Println("parse template failed, err:",err)
		return
	}
	name:="小王子"
	//渲染模板
	t.Execute(w,name)

}
func demo1(w http.ResponseWriter, r *http.Request )  {
	t,err := template.ParseFiles("./t.html","./ul.html")	// 模板 t->ul 顺序不能变
	if err != nil{
		fmt.Printf("parse template failed, err: %v\n",err)
		return
	}
	name := "大王子"
	t.Execute(w,name)
}
func main()  {
	http.HandleFunc("/",f1)
	http.HandleFunc("/tmpDemo",demo1)
	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("http server start failed err:",err)
		return
	}
}