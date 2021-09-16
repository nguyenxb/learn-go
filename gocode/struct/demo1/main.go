package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	age   int
	color string
	arr   [4]int
}
type exmple struct{
	Name string
	map1 map[string]string
}

func (cat *Cat) Print() {
	fmt.Println("name=", cat.Name)
	fmt.Println("age=", cat.age)
	fmt.Println("color=", cat.color)
}
func modifyCat(cat Cat) {
	cat.Name = "modify" + cat.Name
}
func main() {
	/*
		张老太养猫，
		名字 小白，年龄 3，颜色，白色
		名字 小花，年龄 12，颜色，花色
		用户输入名字，就显示名字.年龄，颜色

		如果使用 变量，数组，map来存储，那么会出现数据类型的问题
	*/
	//  结构体时值类型的
	var cat1 Cat // 定义时就分配内存空间
	var cat2 Cat
	var cat3 *Cat = new(Cat) // 获取一个Cat类型的指针
	fmt.Println("cat3=", *cat3)
	fmt.Println(cat1)
	cat1 = Cat{"小白", 3, "white", [4]int{1, 2, 3, 4}}
	cat2 = Cat{"小花", 3, "花", [4]int{4, 5, 6, 7}}
	fmt.Println(cat1)
	fmt.Println(cat2)
	fmt.Println("name=", cat1.Name)
	fmt.Println("Age=", cat1.age)
	fmt.Println("color=", cat1.color)
	fmt.Println("函数")
	cat1.Print()

	fmt.Println(cat1)
	modifyCat(cat1) // 没有发生修改
	fmt.Println(cat1)
	
	fmt.Println()
	exmple := {"name111","asds","1232123"}
	fmt.Println(exmple)
}
