package main

import (
	"fmt"
	"reflect"
)

// 改变普通数据类型的值
func reflect01(b interface{}) {
	fmt.Println("reflect01:")
	// 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v, rVal type = %T,rVal kind : %v\n", rVal, rVal, rVal.Kind())

	// 改变值 // rVal.Elem() // 返回指针指向的值 就比如 *ptr
	// Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
	rVal.Elem().SetInt(20)

}

// 通过反射修改，num int 的值
// 修改student的值
func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("main:")
	fmt.Println("num=", num)
}
