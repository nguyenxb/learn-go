package main

import "fmt"

type Circle struct {
	radius float64
}

func (cricle Circle) area() float64 {
	return 3.14 * cricle.radius * cricle.radius
}

// (*cricle).radius 等价于 cricle.radius ，
// 底层自动优化，最终go使用的还是(*cricle).radius
func (cricle *Circle) area2() float64 {

	return 3.14 * (*cricle).radius * (*cricle).radius
}

func main() {
	var cricle Circle
	cricle.radius = 9
	res := cricle.area()
	fmt.Println("res=", res)

	var cricle1 Circle
	cricle1.radius = 6
	// (&cricle).area() 等价于 cricle.area()
	// 底层自动优化，最终go使用的还是(&cricle).area()
	res1 := cricle1.area()
	// res2 := (&cricle1).area()

	fmt.Println("res=", res1)
}
