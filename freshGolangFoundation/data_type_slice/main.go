package main

import (
	"fmt"
	"time"
)

func modify1(a int)  {
	a = 10
	return
}
func modify2(a *int)  {
	*a = 10
}
func modifyrun()  {
	a := 5
	b := make(chan int, 1)
	fmt.Println("a: ",a)
	fmt.Println("b: ",b)
	modify1(a)
	fmt.Println("m a: ",a)
	modify2(&a)
	fmt.Println("m2a: ",a)
 }


// 对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n!
func sum(n int) uint64  {
	var s uint64 = 1
	var sum uint64 = 0
	for i:=1;i<n;i++{
		s=s*uint64(i)
		fmt.Println("s:",s)
		fmt.Printf("%d!=%d\n",i,s)
		sum += s
	}
	return sum
}
func sumrun()  {
	var n int
	fmt.Printf("请输入一个数字：")
	fmt.Scanf("%d",&n)
	s:=sum(n)
	fmt.Println(s)
}
// 数组
func test1()  {
	var a [10]int
	a[0] = 10
	a[7] = 100
	fmt.Println("a:",a)
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}
	for index,val:=range a{
		fmt.Printf("a[%d]=%d\n",index,val)
	}
}
func test2()  {
	var a [10]int
	b:=a
	b[0]=100
	fmt.Println("a:",a)
	fmt.Println("b:",b)
}
func test3(arr *[5]int)  {
	(*arr)[0] = 1000
}
// 多维数组
func test5()  {
	var f [2][3]int = [...][3]int{{1,2,3},{7,8,9}}
	for k,v := range f{
		for i,j := range v{
			fmt.Printf("(%d,%d)=%d\n",k,i,j)
		}
		fmt.Println()
	}
}
func fab(n int)  {
	var a []int
	a= make([]int,n)
	a[0]=1
	a[1]=1
	for i := 2;i<n;i++{
		a[i] = a[i-1]+a[i-2]
	}
	fmt.Println("a= ",a)
	for _,v:=range a{
		fmt.Println("fab:",v)
	}

}
func testArray()  {
	var a [5]int = [5]int{1,2,3,4,5}
	var a1 = [5]int{1,2,3,4,5}
	var a2 = [...]int{38,78,28,38,348,235,457}
	var a3 = [...]int{1:100,3:200}
	var a4 = [...]string{1:"hello",3:"world"}
	fmt.Println("a:",a)
	fmt.Println("a1:",a1)
	fmt.Println("a2:",a2)
	fmt.Println("a3:",a3)
	fmt.Println("a4:",a4)
}
func testArray2()  {
	var a [2][5]int = [...][5]int{{1,2,3,4,5},{6,7,8,8,9}}
	for row,v := range a{
		for vol,val:=range v{
			fmt.Println("rcv:",row,vol,val)
		}
		fmt.Println()
	}
}
func test6()  {
	const (
		Man = 1
		Female = 2
	)
	currSecond := time.Now().Unix()
	if currSecond%Female==0{
		fmt.Println(Female)
	}else {
		fmt.Println(Man)
	}
}

// 函数传参：做了一个拷贝，如果参数为变量值，则函数内部对变量做的操作不会影响实际变量的值。若传的参数为地址，则对地址指向变量的操作会影响实际变量的值。
func test7(a int)  {
	a = 10
	return
}
func test8(a *int)  {
	*a = 10
}
func modify()  {
	a := 5
	b := make(chan int, 1)
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	test7(a)	//拷贝不影响原来的值
	fmt.Println("a1: ",a)
	test8(&a)
	fmt.Println("a2: ",a)
	// 栈是在栈内存中分配的，公用的，性能最高， 堆是在系统内存中分配的
	// 基本数据类型是在栈中存储的 引用数据类型是在堆中存储的

}
func shuixianhuaDeno(n int) bool {
	var i,j,k int
	i=n%10
	j=(n/10)%10
	k=(n/100)%10
	sum :=i*i*i+j*j*j+k*k*k
	return sum == n
}
//func shuixinahua2Demo(n,int) bool {
//	var result = 0
//	str := strconv.Itoa(n)
//	for i:=0;i<len(str);i++{
//		num := int(str[i]-'0')
//		result+=(num*num*num)
//	}
//	number,err := strconv.Atoi(str)
//	if err != nil{
//		fmt.Println("can not convert %s to int\n",str)
//	}
//	if result == number{
//		return true
//	}
//	return false
//}
func shuixinahuaRun()  {
	var n int
	var m int
	n=100
	m=999
	//fmt.Printf("请输入两个数字：")
	//fmt.Scanf("%d,%d",&n,&m)
	for i:=n;i<=m;i++{
		if shuixianhuaDeno(i) == true{
			fmt.Println(i)
		}
	}
}
func main()  {

	//sumrun()
	//modifyrun()


	//test5()
	//test6()
	//fab(7)
	//testArray()
	//testArray2()
	//modify()
	shuixinahuaRun()

	//test1()
	//test2()
	//var a [5]int
	//test3(&a)
	//fmt.Println(a)

}