// account
package account

import (
	"fmt"
)

type account struct {
	key  string
	loop bool
	// 账户余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对detail 进行拼接处理即可
	detail string
	// "收支\t账号金额\t收支金额\t说  明"
	// 定义一个变量，判断是否有收支行为
	flag bool
}

func NewAccount() *account {
	return &account{
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0,
		note:    "",
		detail:  "收支\t账号金额\t收支金额\t说  明",
		flag:    false,
	}
}

func (this *account) showDetail() {
	fmt.Println("\n------当前收支明细记录-------")
	if this.flag {
		fmt.Println(this.detail)
	} else {
		fmt.Println("目前没有收支行为。。")
	}
}
func (this *account) income() {
	fmt.Println("\n本次收入金额：")
	fmt.Scanln(&this.money)
	if this.money <= 0 {
		fmt.Println("收入的金额不正确,自动返回主菜单")
		return
	}
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}
func (this *account) pay() {
	fmt.Println("\n本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("支出金额不足，自动返回主菜单")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}
func (this *account) isContinue() bool {
	return this.loop
}
func (this *account) exit() {
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
		this.loop = false
	}
}
func (this *account) MainMenu() {
	for this.isContinue() {
		fmt.Println("\n-------家庭收支记账软件-------")
		fmt.Println("       1. 收支明细")
		fmt.Println("       2. 登记收入")
		fmt.Println("       3. 登记支出")
		fmt.Println("       4. 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetail()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("没有这个选项")
		}
	}
	fmt.Println("你退出了家庭收支记账软件")
}
