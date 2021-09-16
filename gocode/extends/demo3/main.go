package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) Sayok() {
	fmt.Println("A Sayok", a.Name)
}
func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	Name string
	A
}

func (b *B) Sayok() {
	fmt.Println("B Sayok", b.Name)
}

func main() {
	/*
		go语言继承的特性：
		1, 继承的结构体可以使用嵌套匿名结构体所有的字段和方法
		即：首字母大写或小写的字段方法都可以使用
		2.调用匿名结构体的字段或方法可以简化比如 b.A.Name 等价于 b.Name
		3.当我们通过 b 访问匿名结构体 A 中的字段或方法时，优先就近原则，先找 b 中的
		字段或者方法，如果没有找到就去找A中的字段或者方法
		4, 如果B和A有相同的字段和方法，就要通过匿名结构体来区分要访问的是
		哪一个结构体的字段和方法，编译器是采用就近原则访问字段和方法的
	*/
	var b B
	b.Name = "tom"
	b.A.Name = "jack"
	b.age = 9
	fmt.Println("B name=", b.Name, ",,A name", b.A.Name)
	fmt.Println("b.A.age=", b.A.age, ",b.age=", b.age)
	b.Sayok()
	b.A.Sayok()
	b.hello()

}
