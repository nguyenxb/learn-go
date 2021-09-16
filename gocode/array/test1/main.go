package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 将A-Z放置到数组中并输出
	var arr [26]byte
	for i := 0; i < len(arr); i++ {
		arr[i] = 'A' + byte(i)
	}

	for _, value := range arr {
		fmt.Printf("%c ", value)
	}
	// 请输出一个数组的最大值，并得到对应的下标
	var intArr = [...]int{1, -1, 9, 90, 11}
	var maxIndex int = 0
	var maxValue int = intArr[0]
	for i := 1; i < len(intArr); i++ {
		if maxValue < intArr[i] {
			maxIndex = i
			maxValue = intArr[i]
		}
	}
	fmt.Println()
	fmt.Println("maxIndex=", maxIndex)
	fmt.Println("maxValue", maxValue)

	var arr2 = [...]int{1, 2, 3, 4, 5, 6, 7, 11, 199}
	sum := 0
	for _, value := range arr2 {
		sum += value
	}
	avg := float64(sum) / 6.0
	fmt.Printf("sum = %d , avg = %f", sum, avg)

	fmt.Println()
	// 随机生成5个数，并将其打印反转
	var arr3 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr3); i++ {
		arr3[i] = rand.Intn(99) + 1
	}
	fmt.Println("arr3=", arr3)
	//	反转数组
	for i := 0; i < len(arr3)/2; i++ {
		arr3[i], arr3[len(arr3)-1-i] = arr3[len(arr3)-1-i], arr3[i]
	}
	fmt.Println(arr3)

	// for i := 0; i < 100000; i++ {
	// 	num := rand.Intn(100) + 1
	// 	fmt.Println("i=", i, "num=", num)
	// 	if num == 100 {
	// 		break
	// 	}
	// }
}
