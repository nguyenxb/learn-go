package main

import (
	"encapsulation/test/model"
	"fmt"
)

func main() {
	account := model.NewAccount("gs1111", "123456", 20)
	fmt.Println(*account)
	account.Login("gs1111", "123456")
	account.Deposite(100)
	account.Query()
	account.WithDraw(120)
	account.Query()
	account.Deposite(50000)
	account.Query()
	account.Logout()
	account.Deposite(10)
}
