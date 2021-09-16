package main

import (
	"fmt"
	// "io"
	"net"
)

func process(conn net.Conn) {
	// 循环接收客户端发送的数据
	defer conn.Close()

	for {
		// 创建一个切片
		buf := make([]byte, 1024)
		// conn.Read(buf)
		// 1.等待客户端通过conn发送信息
		// 2. 如果客户端没有write[发送]，那么协程就会阻塞在这里
		fmt.Println("服务器在等待客户端的输入conn", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn 读取
		if err != nil {
			fmt.Println("服务器的read err=", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("str=-%v-\n", str)
		if str == "exit" {
			break
		}
	}
}
func main() {
	fmt.Println("服务器开始监听。。。")
	// net.Listen("tcp","0.0.0.0:8888")
	// tcp 表示使用网络协议是tcp
	// 0.0.0.0：8888 表示在本地监听8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("lister err", err)
		return
	}
	defer listen.Close() // 延时关闭listen
	for {
		// 等待客户端链接
		fmt.Println("等待客户端来链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {

			fmt.Println("Accep() suc con =", conn)
			//  获取客户端的ip地址
			fmt.Println("客户端ip = ", conn.RemoteAddr())
		}
		// 这里准备起一个协程，为客户端服务
		// 使用协程为连接的客户端服务
		go process(conn)
	}
}
