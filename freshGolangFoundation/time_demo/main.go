package main

import (
	"fmt"
	"time"
)

func timeDemo(){
	now := time.Now()
	fmt.Printf("current time:%v\n",now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)  // %02d保留两位整数，格式化：7=07
}

func timestampDemo()  {
	now := time.Now()		// 当前时间
	timestamp1 := now.Unix()	// 获取当前时间搓
	timestamp2 := now.UnixNano()	// 纳秒时间戳
	fmt.Println("current now: ",now)
	fmt.Println("current timestamp1: ",timestamp1)
	fmt.Println("current timestamp2: ",timestamp2)
}
func timestampDemo2(timestamp int64)  {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)
}

func tickDemo()  {
	ticker := time.Tick(time.Second)	// 定义一个1秒间隔的定时器
	for i := range ticker{
		fmt.Println(i)		// 每秒都会执行任务
	}
}
func formatDemo(){
	//now := time.Now()	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	//fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	//fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) 	// 12小时制
	//
	//fmt.Println(now.Format("2006/01/02 15:04"))
	//fmt.Println(now.Format("15:04 2006/01/02"))
	//fmt.Println(now.Format("2006/01/02"))

	loc,err := time.LoadLocation("Asia/Shanghai") // 加载时区
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("loc:",loc)
	timeStr := "2020/06/08 19:42:00"
	timeObj,err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil{
		fmt.Println(timeObj)
		return
	}
	fmt.Println("timeObj",timeObj)
	timeObj2,err := time.ParseInLocation("2006/01/02 15:04:05", timeStr,loc)  // 按照指定时区和指定格式解析字符串时间
	if err != nil{
		fmt.Println(timeObj2)
		return
	}
	fmt.Println("timeObj2",timeObj2)

}
func main()  {
	//timeDemo()
	//timestampDemo()
	//now := time.Now()		// 当前时间
	//timestamp1 := now.Unix()	// 获取当前时间搓
	//timestampDemo2(timestamp1)
	//n := 5
	////time.Sleep(time.Second*5)
	//time.Sleep(time.Second*time.Duration(n))

	//now := time.Now()
	//later := now.Add(time.Hour)		// now时间+一个小时
	//fmt.Println(later)
	//fmt.Println(later.Sub(now))		// Sub：时间差 ，Equal：相等，Before，After：如果t代表的时间点在u之后，返回真；否则返回假。
	//tickDemo()
	formatDemo()

}