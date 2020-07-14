package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readfromfile()  {
	file,err := os.Open("./main.go")
	if err != nil{
		fmt.Println("open file failed: ",err)
		return

	}
	fmt.Println(file)
	defer file.Close()
	var tmp = make([]byte,128)
	n,err := file.Read(tmp)
	if err == io.EOF{
		fmt.Println("文件读完了")
		return
	}
	if err != nil{
		fmt.Println("failed err:",err)
		return
	}
	fmt.Printf("读取了%d字节数据、\n",n)
	fmt.Println(string(tmp[:n]))
}
func readall()  {
	file,err := os.Open("./xx.log")
	if err != nil{
		fmt.Println("open file failed: ",err)
		return

	}
	fmt.Println(file)
	defer file.Close()
	var content []byte
	var tmp = make([]byte,128)
	for {
	n,err := file.Read(tmp)
	if err == io.EOF{
		fmt.Println(string(tmp[:n]))
		fmt.Println("文件读完了")
		break
	}
	if err != nil{
		fmt.Println("failed err:",err)
		return
	}
	//fmt.Printf("读取了%d字节数据、\n",n)
	content = append(content,tmp[:n]...)
	}
	fmt.Println(string(content))
}

// bufio是在file的基础上封装了一层API，支持更多的功能。
func readbyBufio()  {
	file,err := os.Open("./xx.log")
	if err != nil{
		fmt.Println("failed err:",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line,err := reader.ReadString('\n')
		if err == io.EOF{
			if len(line) !=0{
				fmt.Println(line,"---1")
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil{
			fmt.Println("failed err: ",err)
			return
		}
		fmt.Print(line,"---2")
	}
}

// io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。
func readbyIoutil()  {
	content,err := ioutil.ReadFile("./xx.log")
	if err != nil{
		println("failed err:",err)
		return
	}
	fmt.Println(string(content))
}

func writerbywriter()  {
	file,err := os.OpenFile("xx.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0677)
	if err != nil {
		fmt.Println("failed err：",err)
		return
	}
	defer file.Close()
	str := "xiao小明"
	file.Write([]byte(str))      //写入字节切片数据
	file.WriteString("hello小明") //直接写入字符串数据
}

func writerbybufio()  {
	file,err := os.OpenFile("xx.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0677)
	if err != nil {
		fmt.Println("failed err：",err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file) //直接写入字符串数据
	for i := 0; i<10;i++{
		writer.WriteString("hello 小明") // 将数据写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件,
}
func writerbyioutil() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func catfile(r *bufio.Reader)  {
	for {
		buf,err := r.ReadBytes('\n')
		if err == io.EOF{
			break
		}
		fmt.Fprintf(os.Stdout,"%s",buf)
	}
}

func main()  {
	//readfromfile()
	//readall()
	//readbyBufio()
	//readbyIoutil()
	//writerbybufio()
	//writerbyioutil()
}