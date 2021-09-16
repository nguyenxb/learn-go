package main

import (
	"fmt"
)

//  一个被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}
func main() {
	// 传统的测试方法，就是在main函数中使用看看结果是否正确
	res := addUpper(10) //1 + ... + 10 = 55
	if res != 55 {
		fmt.Printf("不符合期望，期望值%v,结果%v", 55, res)
	} else {
		fmt.Printf("符合期望，期望值%v,结果%v", 55, res)

	}
}
