package main

import (
	"fmt"
	"time"
)

func worker(id int,jobs<-chan int,results chan <- int)  {
	for jobs:=range jobs{
		fmt.Printf("work:%d start job:%d\n",id,jobs)
		results<-jobs*2
		time.Sleep(time.Millisecond*500)
		fmt.Printf("work:%d stop job:%d\n",id,jobs)
	}
}
//func main()  {
//	jobs:=make(chan int,100)
//	results:=make(chan int, 100)
//	// 开启3个goroutine，
//	for j:=0;j<3;j++{
//		go worker(j+1,jobs,results)
//	}
//	//发送5个任务
//	for i := 0;i<5;i++{
//		jobs <- i
//	}
//	close(jobs)
//	for i:=0;i<5;i++{
//		ret:=<-results
//		fmt.Println(ret)
//	}
//}



// select 多路复用 解决：在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。你也许会写出使用遍历的方式来实现获取通道值； Go内置了select
func main()  {
	ch := make(chan int,1)  // 如果多个case同时满足，select会随机选择一个。

	for i:=0;i<10;i++{
		select {
		case x:= <-ch:
			fmt.Println(x)
		case ch<-i:
		default:
			fmt.Println(".......")

		}
	}
}