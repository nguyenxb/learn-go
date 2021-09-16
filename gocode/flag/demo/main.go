package main

import (
	"flag"
	"fmt"
)

func main() {
	// 用flag 包来解析命令行参数
	// 定义几个变量用于接收指定的参数
	var user string
	var pwd string
	var host string
	var port int
	// &user , 就是接收用户命令行输入的 -u 后面的参数值
	// "u" , 就是 -u 指定参数
	// "" 默认值
	// "用户名，默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "host", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3304, "端口号，默认为3304")

	// 转换
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v",
		user, pwd, host, port)
}
