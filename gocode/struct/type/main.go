package main

import (
	"encoding/json"
	"fmt"
)

// 改成小写，就不能调用
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	// 创建一个monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	// 将monster变量序列化为json 字符串
	// json.Marshal()使用到反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误:", err)
	}
	fmt.Println("jsonStr", string(jsonStr))

}
