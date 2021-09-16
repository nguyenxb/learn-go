package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.5:8888")
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	fmt.Println("conn 连接服务器成功=", conn)
	//  功能1 ： 客户端可以发送但单行数据，然后退出
	// os.Stdin 代表标准输入（终端）

	reader := bufio.NewReader(os.Stdin)
	for {

		// 从终端读取一行用户输入,并准备发送个服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}

		// 将line 发送给服务端
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err = ", err)
		}
		fmt.Println("给服务器发送了", n, "个字节")
	}
}
