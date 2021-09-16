package main

import "fmt"

func myfunc3(args ...int) {
	fmt.Println("args:", args)
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Println()
}

func myfunc(args ...int) {
	myfunc3(args...)
	myfunc3(args[1:]...)
	myfunc3(args[2:4]...)
}
func main() {
	myfunc(1, 2, 3, 4, 5)
}
