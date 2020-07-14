package main

import (
	"fmt"
	"strconv"
)


func main()  {
	s1 := "100"
	i1,err := strconv.Atoi(s1)	// Atoi()函数用于将字符串类型的整数转换为int类型. 参数无法转换为int类型，就会返回错误。
	if err != nil{
		println("cont't convert to int")
	}else {
		fmt.Printf("type:%T value:%#v\n",i1,i1)
	}
	i2 := 200
	s2 := strconv.Itoa(i2)		// Itoa()函数用于将int类型数据转换为对应的字符串
	fmt.Printf("type:%T value:%#v\n",s2,s2)

	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-2", 10, 64)
	u, err := strconv.ParseUint("2", 10, 64)

	fmt.Printf("type:%T value:%#v\n",b,b)
	fmt.Printf("type:%T value:%#v\n",f,f)
	fmt.Printf("type:%T value:%#v\n",i,i)
	fmt.Printf("type:%T value:%#v\n",u,u)

}
