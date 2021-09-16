package main

import "fmt"

func main() {
	// 判断一个年份是否为闰年
	fmt.Println("请输入一个年份：")
	var year int
	fmt.Scanln(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "年是闰年")
	} else {
		fmt.Println(year, "年不是闰年")
	}
}
