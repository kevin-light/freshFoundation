package main

import "fmt"

func main() {
	// var a [3]int
	// var b [4]int
	// fmt.Println(a)
	// fmt.Println(b)
	// // 1.定义使用初始值列表的方式初始化
	// var cityarray = [4]string{"北京", "上海", "广州", "深圳"}
	// fmt.Println(cityarray)
	// fmt.Println(cityarray[1])
	// // 编译器推到数组的长度
	// var boolarray = [...]bool{true, false, true, false}
	// fmt.Println(boolarray[0])
	//3.使用索引值的方式初始化
	// var langarray = [...]string{1: "golang", 2: "python", 3: "java"}
	// fmt.Println(langarray)
	// fmt.Printf("type of a: %T\n ", langarray)
	//
	// 1, for循环遍历
	// var cityarray = [4]string{"北京", "上海", "广州", "深圳"}
	// for i := 0; i < len(cityarray); i++ {
	// 	fmt.Println(cityarray[i])
	// }
	// for range遍历
	// for _, value := range cityarray {
	// 	fmt.Println(value)
	// }

	// cityarray := [3][2]string{
	// 	{"北京", "雄安"},
	// 	{"上海", "杭州"},
	// 	{"广州", "深圳"},
	// }
	// fmt.Println(cityarray)
	// fmt.Println(cityarray[2][0])
	// // 二维数组遍历
	// for _,v1 := range cityarray{
	// 	fmt.Println("---0",v1)
	// 	for _,v2 := range v1{
	// 		fmt.Println(v2)
	// 	}
	// }

	//支持的写法
	// a := [...][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	//不支持多维数组的内层使用...
	// b := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// fmt.Println(a)

	// a := [3]int{10, 20, 30}
	// modifyArray(a)

	// fmt.Println(a)
	// b := [3][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// }
	// modifyArray1(b)
	// fmt.Println(b)

	aa := [5]int{1, 3, 5, 7, 8}
	sum := 0
	for _, value := range aa {
		sum += value

		for i, j := range aa {
			for l, k := range aa[i+1:] {
				if j+k == 8 {
					fmt.Printf("(%d, %d) ", i, l+i+1)
				}
			}
		}
	}
	fmt.Println(sum)

}

func modifyArray(x [3]int) {
	x[0] = 100
}
func modifyArray1(x [3][2]int) {
	x[2][0] = 100
}
