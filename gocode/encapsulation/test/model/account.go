package model

import (
	"fmt"
)

type account struct {
	accountID string
	password  string
	balance   float64
	state     bool
}

// 工厂模式函数-构造函数
func NewAccount(id, psw string, ban float64) *account {
	if !(len(id) >= 6 && len(id) <= 10) {
		fmt.Println("账号长度不符合")
		return nil
	}
	if len(psw) != 6 {
		fmt.Println("密码长度不符合")
		return nil
	}
	if ban < 20 {
		fmt.Println("余额数量不符合")
		return nil
	}
	fmt.Println("创建账户成功")
	return &account{
		accountID: id,
		password:  psw,
		balance:   ban,
		state:     false,
	}

}

//登录
func (acc *account) Login(id, psw string) {
	if acc.accountID == id && acc.password == psw {
		acc.state = true
		fmt.Println("登录成功")
		return
	}
	acc.state = false
	fmt.Println("账号或密码不正确,请重新登录")
}

// 退出登录
func (acc *account) Logout() {
	acc.state = false
	fmt.Println("退出登录成功")
}

// 存款
func (acc *account) Deposite(money float64) {
	if acc.state == false {
		fmt.Println("请登录后再存款")
		return
	}
	if money <= 0 {
		fmt.Println("存入金额数量不正确")
		return
	}
	acc.balance += money
	fmt.Println("存款成功")
}

//取款
func (acc *account) WithDraw(money float64) {
	if acc.state == false {
		fmt.Println("请登录后再取款")
		return
	}
	if money <= 0 || money > acc.balance {
		fmt.Println("取款金额不正确")
		return
	}
	acc.balance -= money
	fmt.Println("取款成功")
}

// 查询
func (acc *account) Query() {
	if acc.state == false {
		fmt.Println("请登录后再查询")
		return
	}
	fmt.Printf("你的账号:%v, 余额:%v\n", acc.accountID, acc.balance)
}
