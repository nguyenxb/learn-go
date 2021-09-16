package main

import "fmt"

func main() {
	var num int
	fmt.Print("请输入一个整数：")
	fmt.Scanf("%d", &num)
	switch {
	case num > 0:
		fmt.Println(num, "大于0")
	case num < 0:
		fmt.Println(num, "小于0")
	default:
		fmt.Println(num, "等于0")
	}
}
