package main

import (
	"fmt"
)

// 定义一个结构体
type Account struct {
	AccountID string
	Password  string
	Balance   float64
}

func (account *Account) Deposite(money float64, psw string) {
	// 匹配密码
	if psw != account.Password {
		fmt.Println("你输入的密码不正确")
		return
	}

	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

// 取款
func (account *Account) WIithDraw(money float64, psw string) {
	// 匹配密码
	if psw != account.Password {
		fmt.Println("你输入的密码不正确")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

// 查询余额
func (account *Account) Query(psw string) {
	// 匹配密码
	if psw != account.Password {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("查询成功，你的账号 %v ,余额为%v\n",
		account.AccountID, account.Balance)
}

func main() {
	// 创建对象
	account := &Account{
		AccountID: "gs11111",
		Password:  "666666",
		Balance:   100.0,
	}

	account.Deposite(1500, "666666")
	account.Query("666666")
	account.WIithDraw(593.6, "666666")
	account.Query("666666")
}
