package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn)  {
	defer conn.Close()	// 默认关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err:= reader.Read(buf[:])	// 读取数据
		if err != nil {
			fmt.Println("read from client failed,err", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到clent端发来的数据：", recvStr)
		conn.Write([]byte(recvStr))		// 发送数据
	}
}

func main()  {
	listen, err := net.Listen("tcp","127.0.0.1:20000")	// 开启服务
	if err != nil{
		fmt.Println("listent failed,err:",err)
		return
	}
	for {
		conn, err := listen.Accept()	// 建立连接
		if err != nil{
			fmt.Println("accept failed,err: ", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}