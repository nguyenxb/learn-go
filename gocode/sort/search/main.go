package main

import "fmt"

func binarySearch(arr []int, key int) int {
	var (
		max int = len(arr) - 1
		min int = 0
		mid int
	)
	for min <= max {
		mid = min + (max-min)>>1
		if arr[mid] > key {
			max = mid - 1
		} else if arr[mid] < key {
			min = mid + 1
		} else {
			return mid
		}
	}
	return -1 - mid
}

func quickSort(arr []int, left, right int) {
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

func binary(arr []int, key, left, right int) {
	if left > right {
		return
	}
	mid := left + (right-left)>>1
	if key > mid {
		binary(arr, key, mid+1, right)
	} else if key < mid {
		binary(arr, key, left, mid-1)
	}
	fmt.Println(mid)
}
func main() {
	var arr = [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("\tarr=", arr)
	QuickSort(arr[:])
	fmt.Println("quick\tarr=", arr)
	keyIndex := binarySearch(arr[:], 6)
	fmt.Println("keyIndex=", keyIndex)

	// keyIndex1 := binary(arr[:], 6, 0, len(arr)-1)
	binary(arr[:], 6, 0, len(arr)-1)
	// fmt.Println("keyIndex1=", keyIndex1)

}
