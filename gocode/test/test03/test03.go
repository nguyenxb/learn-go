package main

import (
	"fmt"
	_ "math"
)

func main() {
	// 0到999 中为水仙花数的数
	for num := 0; num < 1000; num++ {
		if num1, num2, num3 := num/100, num%100/10, ((num % 100) % 10); num == num1*num1*num1+num2*num2*num2+num3*num3*num3 {
			fmt.Println(num, "是水仙花数")
		}

	}
}
