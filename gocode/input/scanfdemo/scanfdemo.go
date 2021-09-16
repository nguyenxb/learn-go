package main

import "fmt"

func main() {
	// var name string
	// var age int
	// var sal float32
	// var isPass bool
	// fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)

	// fmt.Println("name=", name)
	// fmt.Println("age=", age)
	// fmt.Println("sal=", sal)
	// fmt.Println("isPass=", isPass)

	// the second method
	var name1 string
	var age1 int
	var sal1 float32
	var isPass1 bool

	fmt.Println("name：")
	fmt.Scanln(&name1)
	fmt.Println("age：")
	fmt.Scanln(&age1)
	fmt.Println("salary：")
	fmt.Scanln(&sal1)
	fmt.Println("ispass：")
	fmt.Scanln(&isPass1)

	fmt.Println("name1=", name1)
	fmt.Println("age1=", age1)
	fmt.Println("sal1=", sal1)
	fmt.Println("isPass1=", isPass1)
}
