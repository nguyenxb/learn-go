package main

import (
	"encapsulation/model"
	"fmt"
)

func main() {
	var person = model.NewPerson("小明")
	person.SetAge(23)
	person.SetWage(5000)
	fmt.Println(*person)
	fmt.Println("姓名：", person.Name, "年龄：", person.GetAge(), "薪资：", person.GetWage())
}
