package main

import (
	"encoding/json" // 序列化数据
	"fmt"
)

/*
	在JS 语言中，一切都是对象，因此，任何的数据类型都可以通过JSON
	来表示，例如字符串，数字，对象，数组，map.结构体

	JSON 键值对时用来保存数据的一种方式

	演示	结构体，map,切片的序列化
*/
type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	// 创建一个对象
	monster := Monster{
		Name:     "牛魔王",
		Age:      123,
		Birthday: "12月3日",
		Sal:      100.0,
		Skill:    "说话",
	}

	// 将monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Printf("%s", data)
}

// 将 map 进行序列化, 序列号后的数据和定义的顺序不一定是一样的
func testMap() {
	var a map[string]interface{}

	// 使用map需要make
	a = make(map[string]interface{}, 3)
	a["张三"] = "asdad"
	a["age"] = 123

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Printf("map序列化：%v\n", string(data[:]))
}

//  对切片[]map[string]interface ，进行序列号
func testSlice() {
	var slice []map[string]interface{}
	slice = make([]map[string]interface{}, 10)
	tmp := make(map[string]interface{})
	tmp["name"] = "张三"
	tmp["age"] = 25
	tmp["address"] = []string{"北京", "上海"}
	slice[0] = tmp
	tmp = make(map[string]interface{})
	tmp["name"] = "李四"
	tmp["age"] = 16
	tmp["other"] = [5]int{1, 2, 4, 5, 6}
	slice[1] = tmp

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Println("slice 序列化：", string(data))
}

//  对基本数据类型进行序列化,意义不大
func testFloat64() {
	var num1 float64 = 2345.456
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Println("slice 序列化：", string(data))
}
func main() {
	testStruct()
	fmt.Println()
	testMap()
	fmt.Println()
	testSlice()
	testFloat64()
	var data [10]byte
	data[0] = 'T'
	data[1] = 'E'
	var str string = string(data[:])

	fmt.Println(str)
}
