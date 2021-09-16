package main

import (
	"fmt"
)

func main() {
	// buildin ：go的内置函数

	// 1 len 用来求长度，如string、array、slice、map、channel
	// func len(v Type) int

	//func new(Type) *Type
	// 内建函数new分配内存。其第一个实参为类型，而非值。
	// 其返回值为指向该类型的新分配的零值的指针。
	/*
		new：主要用来分配内存，主要用来分配值类型，比如int、float32，struct 返回的是指针
		make：用来分配内存，主要用来分配引用类型，比如chan，map，slice
	*/

	num1 := 100
	fmt.Printf("num1 type = %v, num1 value = %v, num1 address = %v\n", num1, num1, &num1)

	num2 := new(int) // *int
	fmt.Printf("num1 type = %v, num1 value = %v, num1 address = %v\n", num2, *num2, &num2)
	*num2 = 100
	fmt.Println("num2=", num2)
}
