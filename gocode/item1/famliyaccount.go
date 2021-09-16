package main

import "fmt"

func main() {
	// 显示主菜单
	key := ""
	loop := true
	// 账户余额
	balance := 10000.0
	// 每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对detail 进行拼接处理即可
	detail := "收支\t账号金额\t收支金额\t说  明"
	// 定义一个变量，判断是否有收支行为
	flag := false
	for loop {
		fmt.Println("\n-------家庭收支记账软件-------")
		fmt.Println("       1. 收支明细")
		fmt.Println("       2. 登记收入")
		fmt.Println("       3. 登记支出")
		fmt.Println("       4. 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("\n------当前收支明细记录-------")
			if flag {
				fmt.Println(detail)
			} else {
				fmt.Println("目前没有收支行为。。")
			}
		case "2":
			fmt.Println("\n本次收入金额：")
			fmt.Scanln(&money)
			if money <= 0 {
				fmt.Println("收入的金额不正确,自动返回主菜单")
				break
			}
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("\n本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("支出金额不足，自动返回主菜单")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("\n是否确定退出？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入：")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("没有这个选项")
		}
	}
	fmt.Println("你退出了家庭收支记账软件")
}
