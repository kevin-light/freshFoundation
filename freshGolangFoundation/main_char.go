package main

import "fmt"

func charDemo()  {
	// byte unit8的别名 ascii吗  rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Printf("c1:%T  c2:%T?", c1, c2)

	s := "hello哈罗"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println("------")
	for _, r := range s {
		fmt.Printf("%c\n", r)
	}
}
func forDemo1()  {
	//for i:=0;i<10;i++{
	//	fmt.Println("i= ",i)
	//}

	//i := 1
	//for i>0{					// 相等 while条件，符合条件一直循环
	//	fmt.Println("i>0")
	//}
	//for true{					// 相等 while条件，符合条件一直循环
	//	fmt.Println("i>0")
	//}
	//for {					// 相等 while条件，符合条件一直循环
	//	fmt.Println("i>0")
	//}

	astr := "hello word,世界"
	for i,v := range astr{
		if i>3{
			continue
		}
		if (i>7){
			break
		}
		fmt.Printf("index[%d] val[%c] len[%d]\n", i, v )
	}
}

func labelDemo()  {
//LABEL1:
//	for i := 0; i <= 5; i++ {
//		for j := 0; j <= 5; j++ {
//			if j == 4 {
//				continue LABEL1
//			}
//			fmt.Printf("i is: %d, and j is: %d\n", i, j)
//		}
//	}
//}
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

func main() {
	//charDemo()
	//forDemo1()
	labelDemo()

}
