package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func unmarshelStruct() {
	// str 在项目开发中是通过网络传输或者读取文件得到的，
	str := "{\"monster_name\":\"牛魔王\",\"age\":123,\"Birthday\":\"12月3日\",\"Sal\":100,\"Skill\":\"说话\"}"

	// 定义一个Monster 实例
	var monster Monster
	// 如果传递的不是指针，序列号的结果就是空值
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmashel err= ", err)
	}
	fmt.Printf("反序列化后 monster=%v\n", monster)
	fmt.Printf("moster.Name=%v,monster.Age=%v,monster.Birthday=%v,monster.Sal=%v,monster.Skill=%v",
		monster.Name, monster.Age, monster.Birthday, monster.Sal, monster.Skill)
}
func unmarshelMap() {
	str := "{\"age\":123,\"张三\":\"asdad\"}"
	var a map[string]interface{}
	// 反序列化
	// 注意： 反序列化map,不需要make,因为make操作被封装到了unmarshel函数中
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println("反序列化 map :", a)
}
func unmarshelSlice() {
	str := "[{\"address\":[\"北京\",\"上海\"],\"age\":25,\"name\":\"张三\"}," +
		"{\"age\":16,\"name\":\"李四\",	\"other\":[1,2,4,5,6]},null,null,null,null,null,null,null,null]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println("slice slice 反序列化：", slice)

}

// 反序列化时数据类型必须一致
func main() {
	unmarshelStruct()
	fmt.Println()
	unmarshelMap()
	fmt.Println()
	unmarshelSlice()
}
