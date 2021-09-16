package main

import "fmt"

func getCount(arr [8]int) (above, less int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Println(avg)
	for i := 0; i < len(arr); i++ {
		if float64(arr[i]) > avg {
			above++
		} else if float64(arr[i]) < avg {
			less++
		}
	}
	return
}

/*
	定义一个数组，给出8个整数，求该数组中大于平均值的数的个数，和小于平均值的数的个数
*/
func main() {
	var arr [8]int = [8]int{5, 2, 3, 4, 5, 6, 7, 8}
	above, less := getCount(arr)
	fmt.Printf("above = %v,less = %v", above, less)
}
