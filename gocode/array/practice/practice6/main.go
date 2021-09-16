package main

import "fmt"

func main() {
	// 查找 字符串数组下是否存在 "AA",
	//如果找到，就输出，如果有多个也要输出完
	var arr [10]string = [10]string{"Abd", "AAb", "AA", "aa", "ANC",
		"aaa", "AA", "AA", "AAA", "AA1"}
	for i := 0; i < len(arr); i++ {
		if arr[i] == "AA" {
			fmt.Printf("在%v位置有\\\"AA\\\"\n", i)
		}
	}
}
