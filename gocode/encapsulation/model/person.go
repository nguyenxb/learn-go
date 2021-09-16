package model

import (
	"fmt"
)

type person struct {
	Name string
	age  int
	wage float64
}

//工厂模式函数，构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问 age 和 wage
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetWage(wage float64) {
	if wage > 0 {
		p.wage = wage
	} else {
		fmt.Println("薪资范围不正确")
	}
}
func (p *person) GetWage() float64 {
	return p.wage
}
