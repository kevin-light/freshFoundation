package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup
func worker(id int,jobs<-chan int,results chan <- int)  {
	for j := range jobs{
		fmt.Printf("worker:%d start job:%d\n",id,j)
		time.Sleep(time.Second)
		//wg.Done()
		fmt.Printf("worker:%d end job %d\n",id,j)
		results <- j*2
	}
}
func wrun()  {
	jobs := make(chan int, 100)
	results := make(chan int,100)
	// 开启3个gotoutine
	for w:=1;w<=3;w++{
		//wg.Add(1)	// 启动一个goroutine就登记 + 1
		go worker(w,jobs,results)
	}
	// 5个任务
	for j:=0;j<7;j++{
		jobs<-j
	}
	close(jobs)
	// 输出结果
	for a:=1;a<=5;a++{
		<-results
	}

}

// select多路复用: 同时从多个通道接收数据。select关键字，可以同时响应多个通道的操作。
func selectRun()  {
	ch:=make(chan int,1)
	for i:=0;i<10;i++{
		select{
		case x := <-ch:
			fmt.Println(x)
		case ch<-i:
		default:
			fmt.Println(".......")
			
			
		}
	}
}

func main()  {
	//wrun()
	selectRun()
}