package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		name     string
		password string
	)
	fmt.Println("请输入用户名：")
	fmt.Scanln(&name)
	fmt.Println("请输入密码：")
	fmt.Scanln(&password)

	if strings.EqualFold(name, "张三") && strings.EqualFold(password, "1234") {
		fmt.Println("登陆成功")
	} else {
		fmt.Println("登录失败")
	}
	// it is same as the above
	if name == "张三" && password == "1234" {
		fmt.Println("登陆成功")
	} else {
		fmt.Println("登录失败")
	}

}
