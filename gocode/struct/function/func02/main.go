package main

import "fmt"

type Person struct {
	name string
}

func (person Person) jisuan() {
	sum := 0
	for i := 1; i < 1000; i++ {
		sum += i
	}
	fmt.Println("jisuan sum = ", sum)
}
func (person Person) jisuan2(n int) {
	sum := 0
	for i := 1; i < n; i++ {
		sum += i
	}
	fmt.Println("jisuan2 sum =", sum)
}
func (person Person) getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	var person Person
	person.jisuan()
	person.jisuan2(100)
	sum := person.getSum(10, 12)
	fmt.Println("main() sum = ", sum)
}
