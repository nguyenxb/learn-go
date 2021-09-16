package main

import (
	"fmt"
)

func main() {
	// defien array
	var numArr1 [3]int = [3]int{1, 2, 3}

	fmt.Println("numArr1=", numArr1)
	var numArr2 = [3]int{4, 5, 6}
	fmt.Println("numArr2=", numArr2)

	var numArr3 = [...]int{8, 9, 10}
	fmt.Println("numArr3", numArr3)

	var numArr4 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr4", numArr4)

	numArr3 = numArr1
	fmt.Println("numArr3", numArr3)
}
