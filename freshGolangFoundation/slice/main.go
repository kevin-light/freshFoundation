package main

import "fmt"

// 动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：
// make([]T, size, cap)
// T:切片的元素类型
// size:切片中元素的数量
// cap:切片的容量

func main() {
	// //  基于数组得到切片
	// a := [5]int{5, 6, 7, 8, 9}
	// b := a[1:5]
	// fmt.Println(b)
	// fmt.Printf("%T\n", b)
	// c := b[0 : len(b)-1]
	// fmt.Println(c)
	// fmt.Printf("%T\n", c)
	//make函数构造切片
	// d := make([]int, 5, 10)
	// fmt.Println(d)
	// fmt.Printf("%T\n", d)
	// fmt.Println(len(d))
	// fmt.Printf("s:%v len(d):%v cap(d):%v\n", d, len(d), cap(d))
	// a := [7]int{1, 2, 3, 4, 5, 7}
	// // t := a[1:4:5]
	// // fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
	// for i := 0; i<len(a); i++{
	// 	fmt.Println(i,a[i])
	// }
	// for index,values := range a{
	// 	fmt.Println(index,values)
	// }
	// var s []int
	// s = append(s, 1)
	// fmt.Println(s)
	// s = append(s, 2, 3)
	// s = append(s, 2, 3)
	// fmt.Println(s)
	// s2 := []int{5, 6, 7}
	// s = append(s, s2...)  // 添加另一个切片中的元素（后面加…）
	// fmt.Println(s)
	// // var s []int
	// s = append(s, 123)
	// fmt.Println(s)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a = append(a[6:])
	// a = append(a[:2],a[6:]...)  // 删除操作  添加另一个切片中的元素（后面加…）
	fmt.Println(a)

}
