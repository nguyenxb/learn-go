package main

import (
	"fmt"
	"strconv"
)

// the string to basic type
func main() {
	str1 := "true"
	b, _ := strconv.ParseBool(str1)
	fmt.Printf("b = %v, its type :%T\n", b, b) //b = true, its type :bool

	str2 := "12"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1 = %v\n", n1, n1) //n1 type int64 n1 = 12
	fmt.Printf("n1 type %T n2 = %v\n", n2, n2) //n1 type int n1 = 12

	var str3 string = "123.345"
	var f1 float64
	var f2 float32
	f1, _ = strconv.ParseFloat(str3, 64)
	f2 = float32(f1)
	fmt.Printf("f1 type %T f1 =%v\n", f1, f1)
	fmt.Printf("f1 type %T f2 =%v\n", f2, f2)

	// note:
	var helloStr string = "hello"
	var num int64 = 10
	// assign to default value
	num, _ = strconv.ParseInt(helloStr, 10, 64)
	fmt.Printf("num type %T num = %v\n", num, num)

	//
	var hello1 string = "hello"
	var boo bool = true
	// assignment to default value
	boo, _ = strconv.ParseBool(hello1)
	fmt.Printf("boo type %T boo = %v\n", boo, boo)

}
