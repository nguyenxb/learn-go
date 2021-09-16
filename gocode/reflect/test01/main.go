package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	// 获取 reflect.Valueof
	rVal := reflect.ValueOf(b)
	// 获取其对于的Type,Kind和值
	fmt.Printf("Type=%v, Kind = %v, 值 = %v\n", rVal.Type(), rVal.Kind(), rVal)
	rVal.Elem().SetFloat(9.9)

	iV := rVal.Interface()
	fmt.Printf("type = %T,Type = %v\n", iV, iV)
	f, ok := iV.(float64)
	if ok {
		fmt.Println("f=", f)
	}
}

func main() {
	var v float64 = 1.2
	reflect01(&v)
	fmt.Println("main v = ", v)

	var str string = "tom"
	fs := reflect.ValueOf(&str) // 必须要传入其指针类型
	fs.Elem().SetString("jack")
	fmt.Println("str=", str)
}
