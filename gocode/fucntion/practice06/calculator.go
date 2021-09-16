package main

import "fmt"

func main() {
	fmt.Println("small calculator: ")
	fmt.Println("1.add")
	fmt.Println("2.subtraction")
	fmt.Println("3.multiplication")
	fmt.Println("4.division")
	fmt.Println("0.exit the program")
	var sel int
	fmt.Scanln(&sel)
	if sel == 0 {
		return
	}
	var (
		num1 float64
		num2 float64
	)
	fmt.Println("please input the first integer")
	fmt.Scanln(&num1)
	fmt.Println("please input the secode integer")
	fmt.Scanln(&num2)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("errors：", err)
		}
	}()
	switch sel {
	case 1:
		fmt.Printf("%.0f + %.0f = %.0f", num1, num2, num1+num2)
	case 2:
		fmt.Printf("%.0f - %.0f = %.0f", num1, num2, num1-num2)
	case 3:
		fmt.Printf("%.0f * %.0f = %.0f", num1, num2, num1*num2)
	case 4:
		if num2 == 0 {
			panic("the denominator can' be zeros")
		}
		fmt.Printf("%.2f / %.2f = %.2f", num1, num2, num1/num2)

	}
}
