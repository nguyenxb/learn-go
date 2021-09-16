package main

import (
	"fmt"
)

// 如果一个类型实现了 Stirng()方法，那么在调用fmt.Println时会自动调用String（）方法
type Integer int

type Student struct {
	name string
	age  int
}

// 结构体是值类型
// 传地址比传值更快
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v], Age=[%v]", stu.name, stu.age)
	return str
}

func (integer Integer) print() {
	fmt.Println("integer=", integer)
}

func main() {
	var i Integer = 10
	i.print()
	stu := Student{
		name: "tom",
		age:  20,
	}
	fmt.Println(stu)
	fmt.Println(&stu)

}
