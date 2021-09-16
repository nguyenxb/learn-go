package main

import (
	"fmt"
)

func main() {
	// 管道中可以放map
	var mapChan chan map[string]string
	// 先创建管道的空间
	mapChan = make(chan map[string]string, 10)

	// 创建map 集合
	map1 := make(map[string]string, 5)
	map1["city1"] = "北京"
	map1["city2"] = "上海"
	map1["city3"] = "天津"

	// 创建 map 集合
	map2 := make(map[string]string, 10)
	map2["name1"] = "小明"
	map2["name2"] = "大名"
	map2["name3"] = "张三"
	map2["name4"] = "李四"
	map2["name5"] = "王五"
	map2["name6"] = "赵六"

	// 将map放入到管道中
	mapChan <- map1
	mapChan <- map2
	fmt.Println(mapChan)

	// 取出
	map11 := <-mapChan
	map22 := <-mapChan
	fmt.Println(map11)
	fmt.Println(map22)
}
