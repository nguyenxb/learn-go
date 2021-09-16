package main

import "fmt"

// variable arguments must be put last
func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	fmt.Println("sum1 =", sum(1, 2, 3))
	fmt.Println("sum2 =", sum(1, 2, 3, 4, 5, 6))
	a, b := 1, 2
	fmt.Println("a=", a, "b=", b)
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)
	a, b = b, a
	fmt.Println("a=", a, "b=", b)
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
