package main

import (
	"fmt"
)

//  Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现。
//  map[KeyType]ValueType ; map类型的变量默认初始值为nil，需要使用make()函数来分配内存

// 动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：
// make([]T, size, cap)
// T:切片的元素类型
// size:切片中元素的数量
// cap:切片的容量

func main() {
	// scoreMap := make(map[string]int, 8)
	// scoreMap["张三"] = 90
	// scoreMap["小明"] = 100
	// fmt.Println(scoreMap)
	// fmt.Println(scoreMap["小明"])
	// fmt.Printf("type of a:%T\n", scoreMap)
	// userInfo := map[string]string{
	// 	"usernamer": "巴拉巴",
	// 	"password":  "123434t6",
	// }
	// fmt.Println(userInfo)
	// key,val := userInfo["username"]  //map键值对是否存在判断
	// if val{
	// 	fmt.Println(key)
	// }else {
	// 	fmt.Println(val)
	// }
	// for k,v:= range userInfo{
	// 	fmt.Println(k,v)
	// }
	// delete(map,key)

	// rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	// var scortMap = make(map[string]int, 200)
	// for i := 0; i < 10; i++ {
	// 	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
	// 	value := rand.Intn(100)          // 生成0-99的随机整数
	// 	scortMap[key] = value
	// }
	// fmt.Println(scortMap)
	// var keys = make([]string, 0, 20)
	// for key := range scortMap {
	// 	keys = append(keys, key) //取出map的key存入切片
	// }
	// fmt.Println(keys)
	// sort.Strings(keys)
	// for _, key := range keys {
	// 	fmt.Println(key, scortMap[key])
	// }
	var mapSlice = make([]map[string]string, 10) // 切片map的初始化
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	// mapSlice[0] = make(map[string]string, 10) // 切片
	// mapSlice[0]["name"] = "bala"
	// mapSlice[0]["password"] = "密码"
	// fmt.Println(mapSlice)
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%v value:%v\n", index, value)
	// }
}
