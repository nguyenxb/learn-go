package main

import (
	"fmt"
	"strings"
)

// 闭包是类，函数是操作，n是字段

func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		// str += "A"
		str += string(36)
		fmt.Println("str =", str)
		return n
	}
}

/*
	1.编写一个函数, makeSuffix(suffix string) 可以接收一个文件
	后缀名，并返回一个闭包
	2. 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（如.jpg）,则
	返回 文件名.jpg ,如果有.jpg后缀，则返回原文件名
	3.要求使用闭包方式完成
	4.strings.HasSuffix

	//
	代码说明：
		返回的匿名函数和makeSuffix(siffix string)的suffix变量 组合
	成一个闭包， 因为 返回的函数引用到了suffix这个变量
		如果使用传统的方法，也可以轻松实现这个功能，但是传统方法需要每次都传入
	 后缀名，比如.jpg,而闭包因为可以保留上次引用的某个值，所以我们传入一次就
	可以反复使用

	即 面向对象编程
*/
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 如果该文件名没有指定的后缀（如.jpg）,则
		// 返回 文件名.jpg ,如果有.jpg后缀，则返回原文件名
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func makeSuffix2(suffix, name string) string {
	// 如果该文件名没有指定的后缀（如.jpg）,则
	// 返回 文件名.jpg ,如果有.jpg后缀，则返回原文件名
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name

}
func main() {
	// get the anonymous function
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16

	// get a closure
	fmt.Println("面向对象：")
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))
	fmt.Println("文件名处理后=", f2("winter.jpg"))
	fmt.Println("文件名处理后=", f2("winter.avi"))
	fmt.Println()
	fmt.Println("面向过程：")
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "winter"))
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "winter.jpg"))
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "winter.avi"))

	fmt.Println()
	f2 = makeSuffix(".avi")
	fmt.Println("文件名处理后=", f2("winter"))
	fmt.Println("文件名处理后=", f2("winter.jpg"))
	fmt.Println("文件名处理后=", f2("winter.avi"))

}
