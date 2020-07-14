package main

import "fmt"

//func main() {
//	//var ch1 chan int  // 引用类型，需要初始化后才能使用
//	//ch1 = make(chan int,1)
//	ch1 := make(chan int, 1)	// 带换冲区通道，异步操作
//	//ch1 := make(chan int)	// 无缓冲区报错，同步通道
//	ch1<-10
//	x:=<-ch1
//	fmt.Println(x)
//
//	close(ch1)
//}

//// 小练习：两个goroutine两个channel，生成0-100的数字发送到ch1，从ch1中取出数据计算平方，把结果发送到ch2
//
//// f1 生成0-100的数字发送到ch1
//func f1(ch chan int)  {
//	for i :=0;i<100;i++{
//		ch<-i
//	}
//	close(ch)
//}
////从ch1中取出数据计算平方，把结果发送到ch2
//func f2(ch1 chan int,ch2 chan int)  {
//	// 从通道中取值的方式1
//	for {
//		tmp,ok := <- ch1
//		ch2<-tmp*tmp
//		if !ok{
//			break
//		}
//
//	}
//	close(ch2)
//
//}
//func main()  {
//	ch1 := make(chan int, 100)
//	ch2 := make(chan int, 200)
//	go f1(ch1)
//	go f2(ch1,ch2)
//	// 从通道中取值的方式2
//	for ret := range ch2{
//		fmt.Println(ret)
//	}
//
//}


// 单向通道，限制通道在函数中只能发送或者接收  ; out只能发送
func counter(out chan <- int)  {
	for i :=0; i < 100; i++{
		out<-i
	}
	close(out)
}
//  out只能发送 ; in只能接收值
func squarer(out chan <- int,in <-chan int)  {
	for i := range in{
		out <- i*i
	}
	close(out)
}
func printer(in<-chan int)  {
	for i := range in{
		fmt.Println(i)
	}
}
func main()  {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2,ch1)
	printer(ch2)
}