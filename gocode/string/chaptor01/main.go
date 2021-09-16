package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num11 int32 = 67
	var num111 int64 = 78

	var num22 float32 = 67.234
	var num2 float64 = 23.345

	var b bool = true
	var mychar byte = 'g'
	var str string // 空
	//the first method
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str = %v, its type : %T \n", str, str)

	str = fmt.Sprintf("%d", num11)
	fmt.Printf("str = %v, its type : %T \n", str, str)

	str = fmt.Sprintf("%d", num111)
	fmt.Printf("str = %v, its type : %T \n", str, str)

	str = fmt.Sprintf("%f", num22)
	fmt.Printf("str = %v, its type : %T\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str = %v, its type : %T\n", str, str)

	// %q , it can print with the ""
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str = %v, its type : %T\n", str, str)
	fmt.Printf("str = %q, its type : %T\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str = %v, its type : %T\n", str, str)
	fmt.Printf("str = %q, its type : %T\n", str, str)

	fmt.Println()
	// the second method
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str = %q, its type :%T\n", str, str)

	str = strconv.FormatInt(int64(num11), 10)
	fmt.Printf("str = %q, its type :%T\n", str, str)

	str = strconv.FormatInt(num111, 10)
	fmt.Printf("str = %q, its type :%T\n", str, str)

	str = strconv.FormatFloat(float64(num22), 'f', 10, 32)
	fmt.Printf("str = %q, its type :%T\n", str, str)

	str = strconv.FormatFloat(num2, 'f', 15, 64)
	fmt.Printf("str = %q,its type :%T\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str = %q, its type :%T", str, str)

	str = strconv.Itoa(num1)
	fmt.Printf("str = %q, its type : %T\n", str, str)

	str = strconv.Itoa(int(num11))
	fmt.Printf("str = %q, its type : %T\n", str, str)

	str = strconv.Itoa(int(num111))
	fmt.Printf("str = %q, its type :%T\n", str, str)
}
