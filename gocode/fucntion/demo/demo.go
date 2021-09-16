package main

import "fmt"

func main() {
	// input two number and an operator,then return a return result
	var (
		n1       float64
		n2       float64
		operator string
		result   float64
	)
	fmt.Println("请输入第一个数：")
	fmt.Scanln(&n1)
	fmt.Println("请输入第二个数：")
	fmt.Scanln(&n2)
	fmt.Println("请输入运算符(+,-,*,/)：")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		result = n1 / n2
	default:
		fmt.Println("无效运算符或无效数值")
	}
	fmt.Printf("%f %s %f = %f", n1, operator, n2, result)
}
