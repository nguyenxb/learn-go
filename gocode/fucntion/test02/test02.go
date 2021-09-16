package main

import "fmt"

// 已知f（1） = 3，f(n) = 2*f(n-1)+1
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}
func main() {
	fmt.Println("f(1)=", f(1))
	fmt.Println("f(5)=", f(5))
}
