package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	t,err:=template.ParseFiles("./index.html")
	if err!=nil{
		fmt.Println("parse template failed, err: ", err)
		return
	}
	msg := "这是index页面"
	t.Execute(w,msg)
}

func home(w http.ResponseWriter, r *http.Request)  {
	t,err:=template.ParseFiles("./home.html")
	if err!=nil{
		fmt.Println("parse template failed, err: ", err)
		return
	}
	msg := "这是 home 页面"
	t.Execute(w,msg)
}

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
	//_,err := template.New("test").Delims("{[", "]}").ParseFiles("./t.tmpl")  // 前端框架（如Vue和 AngularJS）也使用{{和}}作为标识符，同时使用Go语言模板引擎和以上前端框架时就会出现冲突，修改前端的或者修改Go语言的, 这里演示如何修改Go语言模板引擎默认的标识符：

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
func xss(w http.ResponseWriter, r *http.Request){
	tmpl,err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string)template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	str1 := "<script>alert('嘿嘿嘿')</script>"
	str2 := "<a href='http://baidu.com'>百度</a>"
	tmpl.Execute(w,map[string]string{
		"str1":str1,
		"str2":str2,
	})
	//err = tmpl.Execute(w, jsStr)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
func main()  {
	http.HandleFunc("/",f1)
	http.HandleFunc("/tmpDemo",demo1)
	http.HandleFunc("/index",index)
	http.HandleFunc("/home",home)
	http.HandleFunc("/xss",xss)
	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("http server start failed err:",err)
		return
	}
}