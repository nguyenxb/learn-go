package main

import "fmt"

func main() {
	// define an array
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// creating an array slices based on an existing array
	mySlice := myArray[3:7]

	fmt.Println("Element of myArray:")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElement of mySlice1:")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElement of mySlice2:")
	mySlice = append(mySlice, 1, 2)
	fmt.Print(mySlice)
}
