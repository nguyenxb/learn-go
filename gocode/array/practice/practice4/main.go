package main

import (
	"fmt"
)

func main() {
	/*
		定义4行4列的一个数组，从键盘输入值
		对数组的第1，4行数据进行交换位置，
		对数组的第2，3行数据进行交换位置
	*/
	var arr [4][4]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入 %v行%v列 的数：\n", i, j)
			fmt.Scanln(&arr[i][j])
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i < len(arr)/2 {
				// fmt.Printf("%v,%v ", arr[i][j], arr[len(arr)-1-i][j])
				arr[i][j], arr[len(arr)-1-i][j] = arr[len(arr)-1-i][j], arr[i][j]
			}
		}
		// fmt.Println()
	}
	fmt.Println()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

}
