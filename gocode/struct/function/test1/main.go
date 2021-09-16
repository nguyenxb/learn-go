package main

import "fmt"

type MethodUtil struct {
	// 字段。。
}

func (methodUtil MethodUtil) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func (methodUtil MethodUtil) Print2(m, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
func (methodUtil MethodUtil) area(length, width float64) float64 {
	return length * width
}

func (methodUtil *MethodUtil) jugdeNum(num int) bool {
	if num%2 == 0 {
		return false
	} else {
		return true
	}
}
func (methodUtil *MethodUtil) print3(m, n int, key string) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (calculator *Calculator) operation(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calculator.Num1 + calculator.Num2
	case '-':
		res = calculator.Num1 - calculator.Num2
	case '*':
		res = calculator.Num1 * calculator.Num2
	case '/':
		res = calculator.Num1 / calculator.Num2
	default:
		fmt.Println("无效字符")
	}
	return res
}
func (methodUtil *MethodUtil) Print4(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func (methodUtil *MethodUtil) transpose(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i])/2; j++ {
			if i != j {
				arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
			}
		}
	}
}

func main() {
	/*
		1，	编写结构体MethodUtil，编写一个方法，输出一个10*8的矩形
	*/
	var mu MethodUtil
	mu.Print()

	// 2 编写一个方法，提供m和n两个参数，方法中打印m*n的矩形
	fmt.Println("222")
	mu.Print2(3, 2)

	// 3 编写一个方法，提供length，宽width,返回一个面积
	fmt.Println("33333")
	fmt.Println("area=", mu.area(3.5, 4.6))

	// 4 判断一个数是奇数还是偶数
	fmt.Println("444444")
	var num int = 7
	fmt.Println("num=", num, "是否为奇数", (&mu).jugdeNum(num))

	// 5 根据 行m，列n，打印字符如char = "*"
	fmt.Println("555555555")
	mu.print3(3, 4, "#")

	// 6 定义一个结构体 Calcuator,实现+，-，*，/ 四个功能
	var calculator Calculator
	calculator.Num1 = 1.4
	calculator.Num2 = 3.6
	fmt.Println("res add = ", calculator.operation('+'))
	fmt.Println("res sub = ", calculator.operation('-'))
	fmt.Println("res mul = ", calculator.operation('*'))
	fmt.Println("res div = ", calculator.operation('/'))

	// 7 输出乘法表
	fmt.Println("77777777")
	(&mu).Print4(9)
	fmt.Println()
	mu.Print4(5)

	// 8 对给定二维数组进行转置
	fmt.Println()
	fmt.Println("888888")
	var arr [][]int = make([][]int, 3)
	arr[0] = []int{1, 2, 3}
	arr[1] = []int{4, 5, 6}
	arr[2] = []int{7, 8, 9}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
	mu.transpose(arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

}
