package main

import "fmt"

// 猴子吃桃，第一天吃一半再多吃一个，
// 之后每天都吃一半再多吃一个，第十天还有一个桃子
func monkey(n int) int {
	if n == 10 {
		return 1
	} else {
		return (monkey(n+1) + 1) * 2
	}
}

func main() {
	fmt.Println("第1天的桃子数量", monkey(1))
	fmt.Println("第2天的桃子数量", monkey(2))
	fmt.Println("第3天的桃子数量", monkey(3))
	fmt.Println("第4天的桃子数量", monkey(4))
	fmt.Println("第5天的桃子数量", monkey(5))
	fmt.Println("第6天的桃子数量", monkey(6))
	fmt.Println("第7天的桃子数量", monkey(7))
	fmt.Println("第8天的桃子数量", monkey(8))
	fmt.Println("第9天的桃子数量", monkey(9))
	fmt.Println("第10天的桃子数量", monkey(10))
}
