package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)

	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))

	mySlice = append(mySlice, 1, 2)
	fmt.Println("mySlice:", mySlice)
	

	mySlice2 := []int{9, 8, 7, 2, 1}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println("mySlice:", mySlice)

	mySlice = append(mySlice, []int{9, 8, 7, 2, 1}...)
	fmt.Println("mySlice:", mySlice)

	oldSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("cap(oldSlice):", cap(oldSlice))
	fmt.Println()

	newSlice1 := oldSlice[:3]
	fmt.Println("cap(newSlice1):", cap(newSlice1))
	fmt.Println("oldSlice:", oldSlice)
	fmt.Println("newSlice1:", newSlice1)

	fmt.Println()
	newSlice2 := oldSlice[3:]
	fmt.Println("cap(newSlice2):", cap(newSlice2))
	fmt.Println("oldSlice:", oldSlice)
	fmt.Println("newSlice2:", newSlice2)

	fmt.Println()
	newSlice3 := oldSlice[:]
	fmt.Println("cap(newSlice3):", cap(newSlice3))
	fmt.Println("oldSlice:", oldSlice)
	fmt.Println("newSlice3:", newSlice3)

	// fmt.Println()
	// newSlice4 := oldSlice[:10]
	// fmt.Println("cap(newSlice4):", cap(newSlice4))
	// fmt.Println("oldSlice:", oldSlice)
	// fmt.Println("newSlice4:", newSlice4)

	Slice1 := []int{1, 2, 3, 4, 5}
	Slice2 := []int{5, 4, 3}

	fmt.Println()

	copy(Slice2, Slice1) // only the first three elements of Slice1 are copied into Slice2
	fmt.Println("Slice1:", Slice1)
	fmt.Println("Slice2:", Slice2)

	fmt.Println()

	Slice3 := []int{1, 2, 3, 4, 5}
	Slice4 := []int{5, 4, 3}

	copy(Slice3, Slice4) // only the first three elements of Slice1 are copied into Slice2
	fmt.Println("Slice3:", Slice3)
	fmt.Println("Slice4:", Slice4)
}
