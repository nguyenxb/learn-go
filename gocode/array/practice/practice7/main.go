package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr *[10]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func binarySearch(arr [10]int, key int) int {
	max := len(arr) - 1
	min := 0
	var mid int
	for min <= max {
		mid = min + (max-min)/2
		if key > arr[mid] {
			min = mid + 1
		} else if key < arr[mid] {
			max = mid - 1
		} else {
			return mid
		}
	}
	return -1 - mid
}

func main() {
	/*
		随机生成10个整数，（1-100之间）使用冒泡排序法进行排序，然后使用二分
		查找法，查找一个随机的key，
	*/
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	var key int = rand.Intn(100) + 1
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}
	fmt.Println("key=", key)
	fmt.Println(arr)
	bubbleSort(&arr)
	fmt.Println(arr)
	index := binarySearch(arr, key)
	fmt.Println("index=", index)
	if index >= 0 {
		fmt.Printf("数组中有%v这个数,在下标%v位置\n", key, index)
	} else {
		fmt.Println("没有这个数")
	}
}
