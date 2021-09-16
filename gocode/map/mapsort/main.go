package main

import (
	"fmt"
	"sort"
)

func main() {
	// map 排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)

	// 如果按照map的key的顺序进行排序输出
	// 1. 先将map的key放入到切片中
	// 2. 对切片进行排序
	// 3. 遍历切片，然后按照key来输出map的值
	var keys []int
	for key, _ := range map1 {
		keys = append(keys, key)
	}
	// 排序
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)

	for index, key := range keys {
		fmt.Printf("keys[%v] : map[%v]=%v\n", index, key, map1[key])
	}
}
