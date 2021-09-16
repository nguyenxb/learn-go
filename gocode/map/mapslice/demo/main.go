package main

import (
	"fmt"
)

func main() {
	/*
		map slice ：
		使用map切片来存储
	*/
	// 声明一个map切片
	var mosters []map[string]string
	mosters = make([]map[string]string, 2)

	if mosters[0] == nil {
		mosters[0] = make(map[string]string, 2)
		mosters[0]["name"] = "牛魔王"
		mosters[0]["age"] = "500"
	}

	if mosters[1] == nil {
		mosters[1] = make(map[string]string, 2)
		mosters[1]["name"] = "狐狸精"
		mosters[1]["age"] = "400"
	}

	// error: index out of range [2] with length 2
	// if mosters[2] == nil {
	// 	mosters[2] = make(map[string]string, 2)
	// 	mosters[2]["name"] = "玉兔精"
	// 	mosters[2]["age"] = "300"
	// }

	newMoster := map[string]string{
		"name": "新妖怪-火影",
		"age":  "610",
	}
	mosters = append(mosters, newMoster)

	fmt.Println(mosters)

}
