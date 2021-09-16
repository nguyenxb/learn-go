package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "小明"
	p1.Age = 10

	var p2 *Person = &p1

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	p2.Name = "tom-"
	fmt.Printf("p2.Name=%v , p1.Name=%v", p2.Name, p1.Name)
	fmt.Printf("(*p2).Name=%v , p1.Name=%v\n", (*p2).Name, p1.Name)
	fmt.Printf("p2 address = %p\n", p2)
	fmt.Printf("p1 address = %p", &p1)
}
