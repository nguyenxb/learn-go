package main

import "fmt"

// define an global anonymous function
var Func1 = func(n1, n2 int) int {
	return n1 + n2
}

func main() {
	// define the anonymous function and run it directly
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1=", res1)
	// define the anonymous function and assign for an variable
	a := func(n1, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 30)
	fmt.Println("res2=", res2)

	res3 := Func1(10, 4)
	fmt.Println("res3=", res3)
}
