package main

import (
	"fmt"
	"sync"
	"time"
)

//// 多个goroutine并发操作全局变量x,多个add（）函数调用出现数据不准确，
//var(
//	x int64
//	wg sync.WaitGroup
//	lock sync.Mutex   // 互斥锁(浪费时间),它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁
//)
//
//func add()  {
//	for i:=0;i<5000;i++{
//		lock.Lock()
//		x = x+1
//		lock.Unlock()
//	}
//	wg.Done()
//
//}
//func main()  {
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}


// 读写互斥锁
var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex		// 互斥锁
)

func read()  {
	//lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}
func write()  {
	//lock.Lock()
	rwlock.Lock()
	x=x+1
	time.Sleep(time.Millisecond*10)
	rwlock.Unlock()
	wg.Done()
}
func main()  {
	start := time.Now()
	for i:=0;i<1000;i++{
		wg.Add(1)
		go read()
	}
	for i:=0;i<10;i++{
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}