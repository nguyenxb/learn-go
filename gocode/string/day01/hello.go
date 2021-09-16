// hello
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	const (
		v0 = iota
		v1
		v2
	)
	fmt.Println(v0, v1, v2)
	fmt.Printf("%d %d %d", v0, v1, v2)
	fmt.Println()
	fmt.Println("--------------")
	const (
		i          = iota
		i1, i2, i3 = iota, iota, iota
		j1, j2     = iota, iota
	)
	fmt.Println("i = ", i)
	fmt.Println("i1=", i1, "i2=", i2, "i3=", i3)
	fmt.Println("j1=", j1, "j2=", j2)
	var b bool
	fmt.Println("b=", b)
	var aa rune
	fmt.Println("aa=", aa)
	var com complex128
	fmt.Println("com=", com)
	var uintp uintptr
	fmt.Println("unitp=", uintp)
	var s string
	fmt.Println("s=", s)

	var err error
	fmt.Println("err=", err)

	b2 := (1 == 2)
	fmt.Println("b2=", b2)

	str := "Hello, 世界"
	n := len(str)
	fmt.Printf("str = ---%s---\n n = %d\n", str, n)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	fmt.Println("----------------")

	for i, ch := range str {
		fmt.Println(i, ch) // ch 的类型是rune
	}
	// var array []int = []int{1, 2, 3}
	// array := [4]int{1, 2, 3}
	var array [3]int
	array = [3]int{1, 2, 3}
	array = [3]int{3, 4, 5}
	for i, v := range array {
		fmt.Println("Array element [", i, "]=", v)
	}

	var aa1 []int
	aa1 = []int{1, 2, 3, 4, 5}
	fmt.Println(aa1)
	aa1 = []int{1, 2, 3}
	fmt.Println(aa1)

}
