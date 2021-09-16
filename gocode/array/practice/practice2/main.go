package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int, left, right int) {
	// 定义中心轴
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
	}
	arr[p] = pivot
	if p-left > 1 {
		quickSort(arr, left, p-1)
	}
	if right-p > 1 {
		quickSort(arr, p+1, right)
	}
}
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func binarySearch(arr [10]int, key int) int {
	var mid int
	max := len(arr) - 1
	min := 0
	for min < max {
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

func insert(arr [10]int, key int) [11]int {
	var newArr [11]int
	keyIndex := binarySearch(arr, key)
	fmt.Println("keyIndex=", keyIndex)
	if keyIndex >= 0 {
		newArr[keyIndex] = key
	} else if keyIndex < 0 {
		keyIndex = -keyIndex - 1
		if arr[keyIndex] <= key {
			for ; keyIndex < len(newArr)-1; keyIndex++ {
				if keyIndex < len(arr) {
					if arr[keyIndex] > key {
						// fmt.Println("ads")
						break
					}
				}
				// fmt.Println(keyIndex)
			}
		} else {
			for ; keyIndex > 0; keyIndex-- {
				if keyIndex > 0 {
					if arr[keyIndex] < key {
						// fmt.Println("12ads")
						break
					}
				}
				// fmt.Println(keyIndex)
			}
		}
		newArr[keyIndex] = key
		// fmt.Println(keyIndex)
	}
	// fmt.Println(keyIndex)
	for i, j := 0, 0; i < len(newArr); i++ {
		if i != keyIndex {
			newArr[i] = arr[j]
			j++
		}
	}
	// for i := 0; i < len(newArr); i++ {
	// 	if i == -(keyIndex+1) || i == keyIndex {
	// 		newArr[i] = key
	// 	} else if i < keyIndex {
	// 		newArr[i] = arr[i]
	// 	} else {
	// 		fmt.Println(arr[i-1])
	// 		newArr[i] = arr[i-1]
	// 	}
	// }
	return newArr
}
func main() {

	/*
		对一个已经排序的数组，插入一个元素，并且保证该数组是升序的
	*/
	rand.Seed(time.Now().UnixNano())
	var (
		arr [10]int
		key = rand.Intn(100) + 1
	)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}
	// key = -1
	fmt.Println("key=", key)
	fmt.Println(arr)
	QuickSort(arr[:])
	fmt.Println(arr)
	newArr := insert(arr, key)
	fmt.Println(arr)
	fmt.Println(newArr)

}
