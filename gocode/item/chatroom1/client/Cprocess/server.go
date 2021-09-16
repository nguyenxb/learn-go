package Cprocess

import "fmt"

// 显示登录界面
type Cserver struct {
}

func ShowMenu() {

	for {
		var key int
		fmt.Println("\t\t\t 1 显示在线用户列表")
		fmt.Println("\t\t\t 2 发送消息")
		fmt.Println("\t\t\t 3 消息列表")
		fmt.Println("\t\t\t 4 退出登录")
		fmt.Println("\t\t\t 请选择(1-4):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("显示用户列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("消息列表")
		case 4:
			fmt.Println("退出登录")
			return
		default:
			fmt.Println("没有该选项")
		}
	}

}
