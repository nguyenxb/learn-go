package main

import (
	"fmt"
	"fucntion/init/utils"
)

//define an global variable
var age = test() // 1

func test() int {
	fmt.Println("test() run..")
	return 90
}

// init() function which is called before the main()
func init() {
	fmt.Println("init() run..") // 2
}
func main() {
	fmt.Println("main() run.. age =", age) // 3
	fmt.Println("Age=", utils.Age, "Name =", utils.Name)
}
