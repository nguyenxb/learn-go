package main

import (
	"fmt"
	"reflect"
)

// 定义一个结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start--")
	fmt.Println(s)
	fmt.Println("---end---")

}

// 显示s 的值
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	// 获取reflect.Type 类型
	typ := reflect.TypeOf(a)
	// 获取 reflect.Value 类别
	val := reflect.ValueOf(a)
	// 获取 reflect.Kind 类别
	kd := val.Kind()
	// 判断类别
	if kd != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}
	// 获取结构体的字段
	num := val.NumField()
	fmt.Printf("struct has %d fileds\n", num)
	// 遍历结构体的字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d : 值为=%v\n", i, val.Field(i))
		//  读取到struct 标签，注意需要通过reflect.Tyoe来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		// 如果字段有标签tag 就 输出
		if tagVal != "" {
			fmt.Printf("Field %d : tag为=%v \n", i, tagVal)
		}
		// 或者方法的数量
		numofMethod := val.NumMethod()
		fmt.Printf("struct has %d method\n", numofMethod)

		// 调用方法
		// 方法的排序是按照函数名的ascii码排序的
		val.Method(1).Call(nil) // 获取第二个方法，并调用它

		// 调用结构体的第一个方法Method(0)
		var params []reflect.Value // 声明了 []reflect.Value
		params = append(params, reflect.ValueOf(10))
		params = append(params, reflect.ValueOf(40))
		res := val.Method(0).Call(params) // 传入的参数是[]reflect.Value 返回 []reflect
		fmt.Println("res===", res)        //res=== [<int Value>]
		fmt.Println("res===", res[0])     //res=== 50
		fmt.Println("res=", res[0].Int()) //res= 50

	}
}

// 通过反射获取结构体的字段，调用结构体的方法，并获取结构体标签的值
func main() {

	// 创建一个实例
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   500,
		Score: 88.2,
	}
	TestStruct(a)
}
