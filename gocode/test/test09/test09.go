package main

import "fmt"

func main() {
	for sum, i := 100000.0, 0; ; i++ {
		if sum > 50000 {
			sum -= sum * 0.05
		} else {
			sum -= 1000
		}
		if sum < 0 {
			fmt.Println("sum=", sum)
			fmt.Println("i=", i)
			break
		}
	}
}
