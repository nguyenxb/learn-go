package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// const shortForm = "2006-Jan-02"
	// t, _ := time.Parse(shortForm, "1990-Feb-01")
	// fmt.Println(t)

	startTime, _ := time.Parse("2006-01-02", "1990-01-01")
	fmt.Println(startTime)

	// mytime, _ := time.Parse("2006-01-02", "2000-01-01")
	// fmt.Println(mytime)

	//  get the number of days apart  (units / day)
	// fmt.Println(-startTime.Sub(mytime) / 24)
	//
	fmt.Println("enter the date(format:2006-01-02)")
	var timeStr string
	fmt.Scanln(&timeStr)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误：", err)
		}
	}()

	intime := strings.Split(timeStr, "-")
	if len(intime) != 3 {
		panic("输入格式错误")
	}
	year, _ := strconv.Atoi(intime[0])
	month, _ := strconv.Atoi(intime[1])
	day, _ := strconv.Atoi(intime[2])
	// fmt.Println(year)
	// fmt.Println(month)
	// fmt.Println(day)
	if year < 1990 {
		panic("输入年份小于1990年")
	} else if month < 1 || month > 12 {
		panic("输入月份不在 1-12 月之间")
	} else {
		switch month {
		case 1, 3, 5, 7, 8, 10, 11:
			if 1 > day || day > 31 {
				str := strconv.Itoa(year) + "年" + strconv.Itoa(month) + "月 有 31 天"
				panic(str)
			}
		case 4, 6, 9, 12:
			if 1 > day || day > 30 {
				str := strconv.Itoa(year) + "年" + strconv.Itoa(month) + "月 有 30 天"
				panic(str)
			}
		case 2:

			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				if 1 > day || day > 29 {
					str := strconv.Itoa(year) + "年" + strconv.Itoa(month) + "月 有 29 天"
					panic(str)
				}
			} else {
				if 1 > day || day > 28 {
					str := strconv.Itoa(year) + "年" + strconv.Itoa(month) + "月 有 28 天"
					panic(str)
				}
			}
		}
	}
	var timeString string = strconv.Itoa(year) + "-"
	if 0 < month && month < 10 {
		timeString += "0" + strconv.Itoa(month) + "-"
	} else {
		timeString += strconv.Itoa(month) + "-"
	}
	if 0 < day && day < 10 {
		timeString += "0" + strconv.Itoa(day)
	} else {
		timeString += strconv.Itoa(day)
	}
	// fmt.Println(timeString)
	// get the number of apart(the units / day)
	inputime, _ := time.Parse("2006-01-02", timeString)
	// fmt.Println(inputime)
	apart := int(-startTime.Sub(inputime)/24/60/60) / 1000000000
	// fmt.Println(apart)
	sel := apart % 5
	// fmt.Println("sel=", sel)
	if sel < 3 {
		fmt.Println("打鱼")
	} else {
		fmt.Println("晒网")
	}
}
