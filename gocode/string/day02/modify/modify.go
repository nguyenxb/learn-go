// modify
package main

import (
	"fmt"
)

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify,array values:", array)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In main, array values:", array)
}

// compile by
// in this way , the first element of array in main of will be modified
// func modify(array []int) {
// 	array[0] = 10
// 	fmt.Println("In modify,array values:", array)
// }

// func main() {
// 	array := []int{1, 2, 3, 4, 5, 6}
// 	modify(array)
// 	fmt.Println("In main, array values:", array)
// }
