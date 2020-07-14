package main

import (
	"fmt"
	"net/http"
)
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

//func main()  {		// get 请求,不带参数
//
//	resp, err := http.Get("https://www.baidu.com/")
//	if err != nil{
//		fmt.Println("get failed, err: ",err)
//		return
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil{
//		fmt.Println("get body failed,err:", err)
//		return
//	}
//	fmt.Print("1---,",string(body))
//}


// GET请求的参数需要使用Go语言内置的net/url这个标准库来处理
func main(){
	//apiUrl := ""
	getHandler()
}

