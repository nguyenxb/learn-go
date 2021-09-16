package main

import "fmt"

func main() {

	fmt.Println("ok1")
	fmt.Println("ok2")
	goto here
	fmt.Println("ok3")
	fmt.Println("ok4")
here:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
