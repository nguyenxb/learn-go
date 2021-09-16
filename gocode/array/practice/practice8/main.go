package main

import "fmt"

func getMaxAndMin(arr [8]int) (int, int, int, int) {
	var (
		max      int = arr[0]
		min      int = arr[0]
		minIndex int
		maxIndex int
	)
	for i := 1; i < len(arr)-1; i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
		if min > arr[i] {
			min = arr[i]
			minIndex = i
		}
	}
	return max, maxIndex, min, minIndex
}
func main() {
	/*
		编写一个函数，可以接收一个数组，该数有5个数，请找出最大的数和最小的数和对应
		的数组下标
	*/
	var arr = [...]int{1, 3, 4, 2, 5, 7, 3, 4}
	max, maxIndex, min, minIndex := getMaxAndMin(arr)
	fmt.Printf("max=%v, maxIndex=%v, min=%v, minIndex=%v ",
		max, maxIndex, min, minIndex)

}
