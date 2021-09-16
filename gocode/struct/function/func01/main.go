package main

import "fmt"

type Person struct {
	Num int
}

// 给类型Person绑定一个方法
func (person Person) test() {
	person.Num = 10
	fmt.Println(person.Num)
}

// 新定义一个类型 Per ，go语言认为Per，Person是两个类型
type Per Person

func main() {
	var person Person
	person.test() // 调用方法
	var per Per
	person = Person(per) // 强转才能用
	person.test()

	fmt.Println(person.Num) // 传递的不是指针，不会发生值改变
}
