package main

import (
	"fmt"
	"strings"
	// "math"
)

func main() {
	//拼接字符串,长度
	s3 := "hello"
	s4 := "你好"
	s5 := s3 + s4
	fmt.Println(s5, "-len=", len(s5))
	s6 := fmt.Sprintf("%s --- %s", s3, s5)
	fmt.Println(s6)
	// 字符串的分割
	s7 := "how do you do"
	fmt.Println(strings.Split(s7, " "))
	fmt.Printf("%T\t", strings.Split(s7, " ")) // %T 判断数据类型
	// 判断是否包含
	fmt.Println(strings.Contains(s7, "do"))
	//判断前缀
	fmt.Println(strings.HasPrefix(s7, "how"))
	// 判断后缀
	fmt.Println(strings.HasSuffix(s7, "how"), "====s7")
	// 判断子串的位置
	fmt.Println(strings.Index(s7, "do"))
	//判断子串最后出现的位置
	fmt.Println(strings.LastIndex(s7, "do"))
	// join
	s8 := []string{"how", "do", "you", "do"}
	fmt.Println(s8)
	fmt.Println(strings.Join(s8, "+"))

	// // 十进制转换二进制
	// a := 10
	// fmt.Println("%b ---0", a) //输出字符串
	// fmt.Printf("%b --1", a)   // 用fmt格式化输出
	// // unit8 = 0-255
	// var age uint8
	// age = 255
	// fmt.Println("%d", age)

	// // 浮点数
	// fmt.Println(math.MaxFloat32)
	// fmt.Println(math.MaxFloat64)
	// //布尔值
	// var b bool
	// fmt.Println(b)
	// b = true
	// fmt.Println(b)
	// //字符串
	// s1 := "hello go"
	// s2 := "你好哦"
	// fmt.Println(s1, s2)
	// //打印win文件路径
	// fmt.Println("c:\\code\\go.exe")
	// fmt.Println(`as
	// huanha换行
	// sd  、\
	// \t
	// `)

}
