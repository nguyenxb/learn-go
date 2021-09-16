package main

import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[10] = 900
}

//定义一个学生的结构体
type Stu struct {
	Name    string
	age     int
	Address string
}

func main() {
	// map 是引用类型，遵循引用类型的传递机制，在一个函数中
	// 接收map，修改后，会直接修改原来的map
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 30

	fmt.Println(map1)
	modify(map1)
	fmt.Println(map1) //	map1[10] = 900

	// map 容量达到后会自动扩容，如上所示
	students := make(map[string]Stu, 10)
	// 创建2个学生
	stu1 := Stu{"tom", 10, "北京"}
	stu2 := Stu{"mary", 18, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	//  遍历各个学生的信息,v 是结构体
	for k, v := range students {
		fmt.Printf("学生的编号是 %v\n", k)
		fmt.Printf("学生的名字是 %v\n", v.Name)
		fmt.Printf("学生的年龄是 %v\n", v.age)
		fmt.Printf("学生的地址是 %v\n", v.Address)
		fmt.Println()
	}
}
