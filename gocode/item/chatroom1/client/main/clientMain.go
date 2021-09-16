package main

// client

import (
	"fmt"
	"gocode/item/chatroom/client/Cprocess"
)

func main() {

	// 用于记录系统的使用状态
	loop := true

	// 1 显示第一级菜单

	var userId int
	var userPwd string
	var userName string
	for loop {
		// 用于记录用户输入的选项
		var key int
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("-----登录-----")
			fmt.Println("请输入账号：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			up := &Cprocess.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("用户注册")
			fmt.Println("请输入账号：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入昵称：")
			fmt.Scanf("%s\n", &userName)
			up := &Cprocess.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("你退出了系统")
			loop = false
		default:
			fmt.Println("没有这个选项")
		}
	}

}
