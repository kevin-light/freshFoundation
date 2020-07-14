package main

import "fmt"

func main() {
	// score := 65
	// if score >= 90 {
	// 	fmt.Println("A")
	// } else if score > 75 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("C")
	// }

	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	score := 100
	if score > 70 {
		fmt.Println("A+")
	} else {
		fmt.Println("C")
	}

	// for i := 0; i < 10; i++{
	// 	fmt.Println(i)
	// }
	//for省略初始语句，必须保留分号
	i := 10
	for i > 5 {
		fmt.Println(i)
		i--
	}
	// 无限循环
	// for {
	// 	fmt.Println("hello'''")
	// }

	for i := 0; i < 5; i++ {
		if i == 3 {
			// continue
			break
		}
		fmt.Println(i)
	}
	finger := 3
	switch finger {
	case 1:
		fmt.Println("111")
	case 3:
		fmt.Println("333")
	case 5, 6, 7:
		fmt.Println("555")
	}
	sss := "abc"
	switch {
	case finger < 3:
		fmt.Println("12")
	case finger > 3 && finger < 7:
		fmt.Println("357")
	case sss == "sss":
		fmt.Println("str")
	default:
		fmt.Println("jjk")
	}

	var cityArray = [5]string{"北京", "上海", "深圳"}
	fmt.Println(cityArray)
	fmt.Println(cityArray[3])
	var langarr = [...]string{1: "py", 2: "go"}
	fmt.Println(langarr)
	fmt.Printf("%T\n", langarr)
	for index, value := range langarr {
		fmt.Println(index, value)
	}
	for _, value := range langarr {
		fmt.Println(value)
	}

	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			// fmt.Printf("%s\t", v2)
			fmt.Println(v2)
		}
		fmt.Println()
	}

}

// go算术运算符
// func main() {

// 	a := 5
// 	b := 2
// 	fmt.Println(a+b, ",", a-b, ",", a/b, ",", a%b)
// 	a--
// 	b++
// 	fmt.Println(a, b, a >= b, a != b)
// 	// 逻辑运算符 &&, || , !
// 	fmt.Println(a > b && 1 == 1, ",", !(a < b))
// 	// 位运算符号
// 	// fmt.Println(a&b, a|b, a^b, a<<b, a>>b)
// 	// 赋值运算符
// 	a += 3
// 	fmt.Println(a)

// }
