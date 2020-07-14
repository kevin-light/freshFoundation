package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{})  {
	// 不知道被调用这个函数时会传进来什么类型的参数变量： 1 通过类型断言 2借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj,obj.Name(),obj.Kind())
	fmt.Printf("%T,%v\n",obj,obj)
}
func reflectValue(x interface{})  {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T",v,v)
	k := v.Kind()	// 拿到对应值的类型种类
	// 如何得到一个传入的时候类型的变量
	switch k {
	case reflect.Float32:
		// 把反射取到的值转换成一个int32类型的变量
		ret := float32(v.Float())
		fmt.Println(ret)
	}
}
func reflectSetValue(x interface{})  {
	v := reflect.ValueOf(x)
	// elem()用来根据指针取对应的值
	k:= v.Elem().Kind()
	// 如何得到一个传入的时候类型的变量
	switch k {
	case reflect.Int32:
		// 把反射取到的值转换成一个int32类型的变量
		//ret := float32(v.Float())
		//fmt.Println(ret)
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.12)
	}
}
type Cat struct  {

}

type Dog struct  {

}

func main()  {
	//var a float32=1.23
	//reflectType(a)
	//var c Cat
	//reflectType(c)
	//var e []int
	//reflectType(e)
	//var f []string
	//reflectType(f)
	//var aa int32 = 100
	//reflectValue(aa)
	var aaa int32 = 100
	reflectSetValue(&aaa)
	fmt.Println(aaa)

}