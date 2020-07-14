package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func hello()  {
	fmt.Println("hello goroutine")
}

var wg sync.WaitGroup //实现goroutine的同步
func hello1(i int)  {
	defer wg.Done()   // goroutine结束就登记-1
	fmt.Println("hello sync goroutine  ",i)
}
func hello1glroutine()  {
	for i :=0;i<10;i++{
		wg.Add(1)	// 启动一个goroutine就登记 + 1
		go hello1(i)
	}
	wg.Wait()
}
func a()  {
	for i:=1;i<10;i++{
		fmt.Println("a: ",i)
	}
}
func b()  {
	for i:=1;i<10;i++{
		fmt.Println("b: ",i)
	}
}
func gomaxprocesDemo()  {
	runtime.GOMAXPROCS(3)
	go a()
	go b()
	time.Sleep(time.Second)
}

func recv(c chan int)  {
	ret := <-c
	fmt.Println("接收成功", ret)
}
// channel 无缓冲通道也被称为同步通道。
func recvMain()  {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}
// channel 有缓冲通道
func hcrecv()  {
	ch1 := make(chan int, 3)
	ch1 <- 10
	fmt.Println("发送成功", ch1)
}
// channel 练习  for range从通道循环取值
func channelprac()  {
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {				// 开启goroutine将0-100的数发送到ch2中
		for i:=0;i<100;i++{
			ch2<-i
		}
		close(ch2)
	}()
	go func() { 	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
		for{
			i,ok := <-ch2 // 通道关闭后再取值ok=false
			if !ok{
				break
			}
			ch3<-i*i
		}
		close(ch3)
	}()
	for ia:=range ch2{	// 在主goroutine中从ch2中接收值打印  // 通道关闭后会退出for range循环
		fmt.Println(ia)
	}
}

//将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。 = 单向通道
func counter(out chan<- int)  {
	for i:=0;i<100;i++{
		out<-i
	}
	close(out)
}
// chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
func squarer(out chan<-int,in<-chan int){
	for i:= range in{
		out <- i*i
	}
	close(out)
}
// <-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。
func printer(in <-chan int)  {
	for i:=range in{
		fmt.Println(i)
	}
}
func runoi()  {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2,ch1)
	printer(ch2)
}

func main()  {
	//go hello()  // main()函数返回的时候该goroutine
	//hello1glroutine()
	//gomaxprocesDemo()
	//recvMain()
	//hcrecv()
	//channelprac()
	runoi()
}