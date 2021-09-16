package main

import "fmt"

// 斐波那契数列
func test(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return test(n-1) + test(n-2)
	}
}

func main() {
	fmt.Println(test(1))
	fmt.Println(test(2))
	fmt.Println(test(3))
	fmt.Println(test(4))
	fmt.Println(test(5))
}
