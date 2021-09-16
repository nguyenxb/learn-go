package main

import "fmt"

func main() {

	// 区别
	var num1 = 10 / 4
	fmt.Printf("num1 type is %T, num1 = %v\n", num1, num1)

	var num2 int64 = 10 / 4
	fmt.Printf("num2 type is %T, num2 = %v\n", num2, num2)

	var num3 float32 = 10 / 4
	fmt.Printf("num3 type is %T, num3 = %v\n", num3, num3)

	var num4 float32 = 10.0 / 4
	fmt.Printf("num4 type is %T, num4 = %v\n", num4, num4)

	var num5 float32 = 10 / 4.0
	fmt.Printf("num5 types is %T, num5 = %v\n", num5, num5)

	// a % b : a - a / b * b
	fmt.Println("10 % 3 = ", 10%3)     // 10 - 10/3 * 3 = 1
	fmt.Println("10 % -3 = ", 10%-3)   // 10 - 10/(-3)*(-3) = 1
	fmt.Println("-10 % 3 = ", -10%3)   // -10 - (-10) / 3 * 3 = -1
	fmt.Println("-10 % -3 = ", -10%-3) // -10 - (-10)/(-3)*(-3)
}
