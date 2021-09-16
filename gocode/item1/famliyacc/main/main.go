package main

import (
	"item1/famliyacc/account"
)

func main() {
	var acc = account.NewAccount()
	// 显示主菜单
	acc.MainMenu()
}
