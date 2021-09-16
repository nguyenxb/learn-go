package main

import (
	_ "errors"
	"fmt"
)

func PrintMonthDay(year int, month int8) {

	switch month {
	case 1, 3, 5, 7, 8, 10, 11:
		fmt.Printf("%d 年 %d 有 %d 天", year, month, 31)
	case 4, 6, 9, 12:
		fmt.Printf("%d 年 %d 有 %d 天", year, month, 30)
	case 2:
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {

			fmt.Printf("%d 年 %d 有 %d 天", year, month, 29)
		} else {
			fmt.Printf("%d 年 %d 有 %d 天", year, month, 28)
		}
	}
}
func main() {
	var (
		year  int
		month int8
	)
	// for {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("输入数据有误:", err)
		}
	}()
	fmt.Print("请输入年份：")
	fmt.Scanln(&year)
	fmt.Println()
	fmt.Print("请输入月份：")
	fmt.Scanln(&month)
	if year <= 0 || month <= 0 {
		panic("年份或月份只能为正整数")
	} else if month > 12 {
		panic("月份不能大于12")
	}
	PrintMonthDay(year, month)
	// }
}
