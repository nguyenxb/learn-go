package main

import "fmt"

type Student struct {
	Name string
}

// 编写一个函数，判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		index++
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数 是 bool 类型，值是 %v \n", index, x)
		case float32:
			fmt.Printf("第%v个参数 是 float32 类型，值是 %v \n", index, x)
		case float64:
			fmt.Printf("第%v个参数 是 float64 类型，值是 %v \n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数 是 整数 类型，值是 %v \n", index, x)
		case string:
			fmt.Printf("第%v个参数 是 string 类型，值是 %v \n", index, x)
		case Student:
			fmt.Printf("第%v个参数 是 Student 类型，值---是 %v \n", index, x)
		default:
			if _, ok := x.(Student); ok {
				fmt.Printf("第%v个参数 是 Student 类型，值是 %v \n", index, x)
			} else if _, ok := x.(*Student); ok {
				fmt.Printf("第%v个参数 是 *Student 类型，值是 %v \n", index, x)
			} else {
				fmt.Printf("第%v个参数类型不确定，值是 %v \n", index, x)
			}
		}
	}
}

func main() {
	var n1 float32 = 1.2
	var n2 float64 = 3.5
	var n3 int = 100
	var n4 int32 = 15
	var n5 int64 = 128
	var n6 string = "string"
	var stu1 = &Student{
		Name: "tom",
	}
	var stu2 = Student{
		Name: "小明",
	}
	TypeJudge(n1, n2, n3, n4, n5, n6, stu1, stu2)
}
