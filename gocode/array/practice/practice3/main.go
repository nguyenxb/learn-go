package main

import (
	"fmt"
)

func arrPrint(arr [3][4]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

/*
	定义3行4列的一个数组，从键盘输入值
	将四周的数据清0
*/
func main() {
	var arr [3][4]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入%v行,%v列的数：\n", i, j)
			fmt.Scanln(&arr[i][j])
		}
	}
	arrPrint(arr)
	fmt.Println()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			}
		}
	}
	arrPrint(arr)
}
