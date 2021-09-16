package main

import "fmt"

// pointer demo
func main() {
	// define an variable of int type
	var i int = 10
	// the variable i address
	fmt.Printf("the variable type %T, i address = %v ,i =\n", &i, &i, i)

	// ptr is a pointer variable ,
	//its type is *int
	// its value is &i
	var ptr *int = &i
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Printf("ptr address = %v", &ptr)
	fmt.Printf("prt point to the value(the value of variable i)= %v \n", *ptr)

	fmt.Println("i = ", i)
	*ptr = 13
	fmt.Println("i = ", i)
}
