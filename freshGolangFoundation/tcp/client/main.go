package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	conn, err:=net.Dial("tcp","127.0.0.1:20000") 	// 与服务端建立连接
	if err != nil{
		fmt.Println("err: ",err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')	// 读取用户输入
		input = strings.TrimSpace(input)				// 去空格
		inputInfo := strings.Trim(input, "\r\n")	// 转换大写
		if strings.ToUpper(inputInfo) == "Q"{			// 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo))		// 发送数据
		if err != nil{
			fmt.Println("recv failed , err: ",err)
			return
		}
		buf := [512]byte{}
		n,err := conn.Read(buf[:])					// 从服务端接收回复的信息
		if err != nil {
			fmt.Println("recv failed,err: ",err)
			return
		}
		fmt.Println("收到服务端回复",string(buf[:n]))
	}
}
