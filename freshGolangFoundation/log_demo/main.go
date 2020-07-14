package main

import (
	"fmt"
	"log"
	"os"
)

const (
	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // 日期：2009/01/23
	Ltime                         // 时间：01:23:23
	Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	LUTC                          // 使用UTC时间
	LstdFlags     = Ldate | Ltime // 标准logger的初始值
)


// demo1
//func main()  {
//	logFile,err := os.OpenFile("./tdemo.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err!= nil{
//		fmt.Println("log failed:",err)
//		return
//	}
//	log.SetOutput(logFile)		//SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。
//	//log.Println("打印普通日志")
//	//v := "普通的"
//	//log.Printf("%s 日志、\n", v)
//	//log.Fatalln("会触发fatal的日志")
//	//log.Panicln("会触发panic的日志")
//
//	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
//	log.Printf("日志log1")
//	log.SetPrefix("[san]") 		// 其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。
//	log.Printf("日志log2")
//}

// demo2
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

//func New(out io.Writer, prefix string, flag int) *Logger

func main() {
	//logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	log.Println("这是自定义的logger记录的日志。")
}