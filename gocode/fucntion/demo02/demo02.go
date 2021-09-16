package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}
func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}
func main() {
	fmt.Println("test function:")
	test(4)
	fmt.Println("----------")
	fmt.Println("test2 function ")
	test2(4)
}
