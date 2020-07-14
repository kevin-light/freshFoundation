package main

import (
	"fmt"
	"sort"
)

// 创建切片
//var slice []type = make([]type, len)
//slice := make([]type,len)
//slice := make([]type,len,cap)
//slice = append(slice 10)
//var a = []int{1,2,3}
//var b = []int{4,5,6}
//a = append(a,b,...)
func creatSlice()  {
	s1 := []int{1,2,3,4,5}

	s2 := make([]int, 3)

	copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(s1)

}

type slice struct {
	ptr *[100]int
	len int
	cap int
}

func make1(s slice,cap int) slice  {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}
func testslice()  {
	var slice []int
	var arr [5]int = [...]int{1,2,3,4,5}
	slice = arr[:]
	fmt.Println(slice)
	fmt.Println(arr[2:4])
	fmt.Println(arr[2:])
	fmt.Println(arr[:1])
	fmt.Println(arr[:len(arr)-1])
}
func modify(s slice)  {
	s.ptr[1] = 1000
}
func testslice2()  {
	var s1 slice
	s1 = make1(s1,10)
	s1.ptr[0]=100
	modify(s1)
	fmt.Println("s1.ptr",s1.ptr)

}
func modify1(a []int)  {
	a[1] = 1000
}
func testslice3()  {
	var b []int = []int{1,2,3,4}
	modify1(b)
	fmt.Println(b)

}
func testslice4()  {
	var a = [10]int{1,2,3,4}
	b := a[1:5]
	fmt.Printf("b= %p\n",b)
	fmt.Printf("&a = %p\n",&a[1])
}

// 实例2
func testslice5()  {
	var a [5]int = [...]int{1,2,3,4,5}
	s:=a[1:]
	fmt.Println("a:",a)
	s[1] = 100
	fmt.Printf("s=%p a[1]=%p\n",s,&a[1])
	fmt.Println("befer a:",a)
	s = append(s,10)
	s = append(s,10)
	s = append(s,10)
	s[1] =1000
	fmt.Println("after a:",a)
	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p\n",s,&a[1])
}
func testcopy1()  {
	var a = []int{1,2,3,4,5}
	b := make([]int,10)
	copy(b,a)
	fmt.Println("a:",a)
	fmt.Println(b)
}
func teststring()  {
	s:="hello world"
	s1 := s[0:5]
	s2 := s[6:]
	fmt.Println(s1)
	fmt.Println(s2)
}
func testModifyStr()  {
	s:="收到hello world"
	s1:=[]rune(s)
	fmt.Println("s1:",s1)
	s1[0] = 200
	s1[1] = 128
	fmt.Println("s1  :",s1)
	s1[2] = 256
	str := string(s1)
	fmt.Println(str)
}
func testIntSort()  {
	var a = [...]int{1,4,2,5,73,5,3,77,789}
	sort.Ints(a[:])
	fmt.Println("sa:",a)
}
func testStrSort()  {
	var a = [...]string{"abc", "efg", "b", "A", "eeee","dd","gaer","cc","aa"}
	sort.Strings(a[:])
	fmt.Println(a)
}

// map
//var map1 map[keytype]valuetype
//var a map[string]string
//var a map[string]int
//var a map[int]string
//var a map[string]map[string]string
func testMap1()  {
	//var a map[int]int
	//a = make(map[int]int, 5)
	var a map[string]string = map[string]string{"hello":"world"}
	fmt.Println("a1:",a)
	a = make(map[string]string,10)
	fmt.Println("a1:",a)
	a["hello"]="world"  // 插入和更新
	//val,ok = a["hello"]    //查找
	//fmt.Println("val:",val)
	for k,v := range a{    //遍历
		fmt.Println("k,v",k,v)

	}
	delete(a,"hello")  // 删除
	fmt.Println(len(a))
}
// map是引用类型
func modifyMap1(a map[string]int)  {
	a["one"] = 1314
	fmt.Println("am:",a)
	items:=make([]map[int]int,5)
	fmt.Println("t1:",items)
	for i:=0;i<len(items);i++{
		items[i] = make(map[int]int)
	}
	fmt.Println("t2",items)
}
func main()  {

	// 切片 排序
	//creatSlice()
	//testslice()
	//testslice2()
	//testslice3()
	// 后边常用
	//testslice4()
	//testslice5()
	//testcopy1()
	//teststring()
	//testModifyStr()
	//testIntSort()
	//testStrSort()

	// map key-value的数据结构，又叫字典或关联数组
	//testMap1()
	//modifyMap1()
}