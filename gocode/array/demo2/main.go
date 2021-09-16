package main

import (
	"fmt"
)

func test1(arr [3]int) {
	arr[0] = 88
}
func test2(arr *[3]int) {
	arr[0] = 99
}
func test3(arr []int) {
	arr[0] = 100
	fmt.Println("test3()", arr)
}
func main() {
	var arr1 [3]int = [3]int{11, 22, 33}
	test1(arr1)
	fmt.Println(arr1)
	test2(&arr1)
	fmt.Println(arr1)
	// test3(&arr1) // 编译错误
}
