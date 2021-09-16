package main

import "fmt"

// 斐波那契(递归)
func fbn(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return fbn(n-1) + fbn(n-2)
}

// 斐波那契数
func fbn1(n int) []uint64 {
	// define a slice ,size is n
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
func main() {
	a := fbn(2)
	fmt.Println(a)
	fbnSlice := fbn1(5)
	fmt.Println(fbnSlice)
}
