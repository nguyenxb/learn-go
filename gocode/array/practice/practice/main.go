package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 快速排序
func quickSort(arr []int, left, right int) {
	// 定义一个中心轴
	pivot := arr[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && arr[j] >= pivot {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}
		if i <= p && arr[i] <= pivot {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
		arr[p] = pivot
		if p-left > 1 {
			quickSort(arr, left, p-1)
		}
		if right-p > 1 {
			quickSort(arr, p+1, right)
		}
	}
}
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

// 反转数组
func reverse(arr []int) {
	for left, right := 0, len(arr)-1; left < right; {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
		// fmt.Println(arr, "left=", left, "rignt=", right)
	}
	// fmt.Println(arr)
}

// 二分查找
func binarySearch(arr []int, key int) int {
	var max = len(arr) - 1
	var min = 0
	var mid int
	for min <= max {
		mid = min + (max-min)>>1
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
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
		// 睡眠100 毫秒
		time.Sleep(100 * time.Millisecond)
	}
	// arr[0] = 10
	fmt.Println(arr)
	QuickSort(arr[:])
	// n := binarySearch(arr[:], 10)
	// fmt.Println(n)
	fmt.Println(arr)
	reverse(arr[:])
	fmt.Println(arr)
	// 反转数组
	// for i := 0; i < len(arr)/2; i++ {
	// 	arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	// }
	// fmt.Println(arr)
	sum := 0
	for _, value := range arr {
		sum += value
	}
	max := arr[len(arr)-1]
	min := arr[0]
	fmt.Printf("sum = %v, max = %v, min = %v\n", float64(sum)/float64(len(arr)), max, min)
	reverse(arr[:])
	fmt.Println(arr)
	// index := binarySearch(arr[:], 55)
	// fmt.Println(index)
	if index := binarySearch(arr[:], 55); index >= 0 {
		fmt.Printf("数组中有55,位置在%v角标下", index)
	} else {
		fmt.Println("数组中没有55")
	}
}
