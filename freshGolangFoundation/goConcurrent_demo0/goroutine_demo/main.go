package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func hello(i int)  {
	fmt.Println("hello ",i)
	wg.Done()
}
//func main()  {		// 开启一个主goroutine去执行main函数
//	wg.Add(100)  // 启动一个goroutine就登记+1
//	for i:=0;i<100;i++{
//		//wg.Add(i)
//		go hello(i)		//
//
//	}
//	fmt.Println("hello main")
//
//	//time.Sleep(time.Second) // 不建议
//	wg.Wait()  // 等待所有登记的goroutine都结束
//}

func main()  {		// 开启一个主goroutine去执行main函数
	wg.Add(100)  // 启动一个goroutine就登记+1
	for i:=0;i<100;i++{
		//wg.Add(i)
		go func(i int) {
			fmt.Println("hello",i)
			wg.Done()
		}(i)		// 并发使用 匿名函数+闭包: 必须传参数i int，

	}
	fmt.Println("hello main")

	//time.Sleep(time.Second) // 不建议
	wg.Wait()  // 等待所有登记的goroutine都结束
}