package main

import "fmt"

func main() {
	num, sum, count := 1, 0, 0
	// count := 0
	for {
		if num%9 == 0 {
			count++
			sum += num
		} else if num >= 100 {
			break
		}
		num++
	}
	fmt.Println("num =", num, "sum =", sum, "count =", count)
}
