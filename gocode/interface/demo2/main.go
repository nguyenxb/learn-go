package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type Integer int

func (i Integer) Say() {
	fmt.Println("Integer Say()")
}

/*
	接口本身不能创建实例，但是可以指向一个实现了该接口的自定义
	类型的变量(实例)

	只要是自定义数据类型，就能实现接口，不是只有结构体才能创建接口

	一个自定义的数据类型可以实现多个接口
*/
func main() {
	var stu Stu
	var a AInterface = stu
	a.Say()
	var inte Integer
	inte.Say()

}
