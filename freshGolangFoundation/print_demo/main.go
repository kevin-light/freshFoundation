package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	//fmt.Print("dhello 小明。")		// 输出到系统的标准输出，区别在于Print函数直接输出内容
	//name := "xiaoming"
	//fmt.Print(name)
	//fmt.Printf("%s\n",name)		// 函数支持格式化输出字符串
	//fmt.Println("打印单独一行") 		// 输出内容的结尾添加一个换行符

	//fmt.Fprintln(os.Stdout,"向标准输出写入内容")  // 将内容输出到一个io.Writer接口类型的变量w中，通常用这个函数往文件中写入内容。
	//fileObj, err := os.OpenFile("./balala.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil{
	//	fmt.Println("打开文件出错：err：", err)
	//	return
	//}
	//name1 := "小明来啦"
	//fmt.Fprintf(fileObj,"往文件中写入信息：%s ; ",name1)

	//s1 := fmt.Sprint("小明")		// Sprint系列函数会把传入的数据生成并返回一个字符串。
	//name := "小红"
	//age := 18
	//s2 := fmt.Sprintf("name:%s, age:%d", name,age)
	//s3 := fmt.Sprintln("小强")
	//fmt.Println(s1,s2,s3)
	//
	//fmt.Printf("%v\n",false)
	//o := struct {
	//	name string
	//}{"小明"}
	//fmt.Printf("%v\n",o) 	// %v	值的默认格式表示
	//fmt.Printf("%T\n",0)	// %T	打印值的类型
	//fmt.Printf("100%%\n")		// %%	百分号

	//var(
	//	name string
	//	age int
	//	married bool
	//)
	//fmt.Scan(&name, &age, &married) 	// Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。  Scanln类似Scan，它使用空格分隔。在遇到换行时才停止扫描。
	//fmt.Printf("扫描结果 name:%s age%d married:%t \n",name,age ,married)

	// 完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}