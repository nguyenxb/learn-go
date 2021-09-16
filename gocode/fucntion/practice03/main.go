package main

import "fmt"

func primeNumber() {
	fmt.Printf("100以内的所有素数:")
	var count int = 0
outside:
	for i := 2; i <= 100; i++ {
		for j := 2; j < i; j++ {
			// fmt.Println(i, ":", j)
			if i%j == 0 {
				continue outside
			}
		}

		if count%5 == 0 {
			fmt.Println()
		}
		fmt.Printf("%v\t", i)
		count++
	}
	fmt.Println()
}

func main() {
	primeNumber()
}
