package main

import "fmt"

type Binterface interface {
	test01()
}
type Cinterface interface {
	test01()
	test02()
	test03()
}
type Ainterface interface {
	Binterface
	Cinterface
	test03()
}
type Stu struct {
}

func (stu *Stu) test01() {
	fmt.Println("Binterface")
}
func (stu Stu) test02() {
	fmt.Println("Cinterface")
}
func (stu Stu) test03() {
	fmt.Println("Ainterface")
}

/*
	接口也支持继承，
	只有实现了A接口的所有方法，才是实现了A接口

	接口是引用类型，任何的方法都可以赋值给空接口
*/
func main() {
	var stu Stu
	stu.test01()
	stu.test02()
	stu.test03()
	// 如果要编译通过 改成	var a Ainterface = &stu
	var a Ainterface = stu
	a.test01()
	a.test02()
	a.test03()

	var b Binterface = stu
	b.test01()
	// b.test02() //  不能使用b接口调用属于其他接口的方法，
	// b.test03()
}
