package main

import (
	"fmt"
)

func main() {
	// 保存 1 3 5 7 9到数组，并倒序打印
	var arr [6]int
	for i := 0; i < len(arr); i++ {
		arr[i] = 2*i + 1
	}
	fmt.Println(arr)
	// reverse
	for i := 0; i < len(arr)/2; i++ {
		fmt.Println(i)
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println(arr)
}
