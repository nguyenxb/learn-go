package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MyPrint(arr [4][6]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	// 定义二维数组
	var arr [4][6]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			arr[i][j] = rand.Intn(100)
		}
	}
	MyPrint(arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[0][0])
	fmt.Printf("%p\n", &arr[0][1])
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[2])
	fmt.Printf("%p\n", &arr[3])

	var arr1 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var arr2 [3][3]int = [...][3]int{{11, 12, 13}, {14, 5, 6}, {7, 8, 9}}
	var arr3 = [3][2]int{{1, 2}, {2, 3}, {3, 4}}
	var arr4 = [...][2]int{{1, 2}, {4, 5}}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println()
	//  遍历
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
	// for range
	for i, v := range arr1 {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v] = %v\t", i, j, v2)
		}
		fmt.Println()
	}

}
