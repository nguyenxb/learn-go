package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取的传入的变量的 type , kind 值
	// 1， 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Printf("reflectTest01  rTyp  = %v,rTyp type = %T \n", rTyp, rTyp) //  int

	// 获取到reflect.Value 类型
	rVal := reflect.ValueOf(b)
	fmt.Printf("reflectTest01  rVal= %v, rVal type= %T\n", rVal, rVal)

	// 进行相加运算
	n2 := 4 + rVal.Int()
	fmt.Printf("reflectTest01  n2=%v, n2 type = %T\n", n2, n2)

	// 下面我们将rVal 转成interface{}
	iV := rVal.Interface()
	fmt.Printf("reflectTest01  iV = %v, iV type = %T\n", iV, iV)
	// 将interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("reflectTest01  num2 = %v,num2 type =%T\n", num2, num2)

}

func reflectTest02(b interface{}) {
	// 获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Printf("reflectTest02 rTyp = %v,rType type = %T\n", rTyp, rTyp)

	// 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("reflectTest02  rVal = %v, rVal type = %T\n", rVal, rVal)

	// 将rVal转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("reflectTest02  iV = %v,iV type = %T\n", iV, iV)

	//  将interface{} 通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Println("断言转化成功")
		fmt.Printf("reflectTest02 stu = %v,stu type = %T\n", stu, stu)
		fmt.Println("stu.Name:", stu.Name)
	} else {
		fmt.Println("断言转化失败")
	}

}

type Student struct {
	Name string
	Age  int
}

// 演示反射
func main() {
	// 演示对(基本数据类型，interface{},reflect.Value)
	// 进行反射的基本操作

	// 1 ,先定义一个 int
	var num int = 100
	fmt.Printf("main() num=%v, num=%T\n", num, num)
	reflectTest01(num)

	fmt.Println()
	//  定义 一个student实例
	student := Student{
		Name: "小明",
		Age:  25,
	}
	fmt.Printf("main() student=%v, student=%T\n", student, student)
	reflectTest02(student)
}
