package main

import "fmt"

func main() {
	var (
		year  int
		month int8
	)
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)

	if 1 <= month && month <= 12 {
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			fmt.Printf("%d 年 %d 月 的天数是 %d", year, month, 31)
		case 2:
			if year%400 == 0 || (year%4 == 0 && year%100 == 0) {
				fmt.Printf("%d 年 %d 月 的天数是 %d", year, month, 29)
			} else {
				fmt.Printf("%d 年 %d 月 的天数是 %d", year, month, 28)
			}
		default:
			fmt.Printf("%d 年 %d 月 的天数是 %d", year, month, 30)
		}
	} else {
		fmt.Println("输入有误")
	}
}
