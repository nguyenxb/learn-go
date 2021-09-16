package main

import "fmt"

func getSum(n1, n2 int) int {
	return n1 + n2
}

// 函数是一种数据类型，在go中，函数可以作为形参，并且调用
func myFun(funvar myFunType, num1, num2 int) int {
	return funvar(num1, num2)
}

// define a type which is func(int,int)int
type myFunType func(int, int) int

func main() {
	// 将函数的地址传给变量a
	a := getSum
	fmt.Printf("a的类型%T,getSum的类型%T", a, getSum)

	result := a(10, 40) // 等价 result := getSum(10,40)
	fmt.Println("result=", result)

	//
	result2 := myFun(getSum, 50, 60)
	fmt.Println("result2=", result2)
}
