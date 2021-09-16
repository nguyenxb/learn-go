package main

import (
	"fmt"
	"struct/factory/model"
	"struct/factory/model1"
)

func main() {
	// 当Student结构体是首字母大写的时候，直接创建就可以
	// 创建Student实例
	var student model.Student = model.Student{
		Name: "小明",
		Age:  18,
	}
	fmt.Println(student)

	/// 当student结构体是首字母小写的时候，我们可以通过工厂模式解决
	var student1 = model1.NewStudent("小红", 89)
	fmt.Printf("%T\n", student1)
	fmt.Println("student1:", *student1)

	// 获取结构体的私有的字段
	fmt.Println("score=", student1.GetScore())
}
