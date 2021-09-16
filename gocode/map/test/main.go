package main

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string,
	name string) {
	// 用户存在
	if users[name] != nil {
		users[name]["password"] = "888888"
	} else {
		// 用户不存在
		users[name] = make(map[string]string, 2)
		users[name]["password"] = "9999"
		users[name]["nickname"] = "昵称-" + name

	}
}

func main() {
	/*
		1、使用map[string]map[string]string 的 map 类型
		2、key : 表示用户名，是唯一的，不能重复
		3、如果某个用户名存在，将其密码改成"888888"，如果不存在
		就增加这个用户信息，（包括昵称nickname）
		4、编写一个函数，modifyUser(users map[string]map[string]string,name string)
		完成上述功能

	*/
	map1 := map[string]string{
		"nickname": "昵称-adsad",
		"password": "9090909",
	}
	map2 := map[string]string{
		"nickname": "昵称-aaaa",
		"password": "4564564",
	}
	var map3 map[string]string
	map3 = make(map[string]string, 2)
	map3["nickname"] = "昵称-pppp"
	map3["password"] = "123123"

	var users map[string]map[string]string
	users = make(map[string]map[string]string)

	users["no1"] = map1
	users["no2"] = map2
	users["no3"] = map3

	fmt.Println(users)
	modifyUser(users, "no1")
	fmt.Println(users)
	modifyUser(users, "no5")
	fmt.Println(users)
}
