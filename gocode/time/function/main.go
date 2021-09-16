package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 时间和日期相关函数和方法使用
	// 1 get the current time
	now := time.Now()
	fmt.Printf("now = %v, now type = %T", now, now)

	// 2 通过now 可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化时间日期
	fmt.Printf("当前年月日 ： %v-%v-%v %d:%d:%d\n", now.Year(),
		int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf(" %v-%v-%v %d:%d:%d\n", now.Year(),
		int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr=", dateStr)

	//格式化时间日期的第二种方式
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format(" 15:04:05"))
	// 时间必须设置为 "2006-01-02 15:04:05"
	// 数字是固定的，否则不正确
	// 错误示例
	//20007-02-20 12:57:11
	fmt.Println(now.Format("2007-01-02 15:04:05"))

	// 时间常量
	/*
	   Nanosecond  Duration = 1 // 纳秒
	   Microsecond          = 1000 * Nanosecond // 微秒
	   Millisecond          = 1000 * Microsecond // 毫秒
	   Second               = 1000 * Millisecond // 秒
	   Minute               = 60 * Second
	   Hour                 = 60 * Minute
	*/
	// 每隔一秒输出一个数
	// i := 1
	// for {
	// 	fmt.Println("i=", i)
	// 	i++
	// 	time.Sleep(time.Second)
	// 	if i == 10 {
	// 		break
	// 	}
	// }
	// // 每隔100 毫秒输出一个数
	// for {
	// 	fmt.Println("i=", i)
	// 	i++
	// 	time.Sleep(time.Millisecond)
	// 	if i == 100 {
	// 		break
	// 	}
	// }

	/// 获取时间戳：
	// Unix 时间戳：func (t Time) Unix() int64
	//Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）
	// Unixnano 时间戳：func (t Time) UnixNano() int64
	// UnixNano将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）。
	//如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话，结果是未定义的。
	fmt.Printf("unix时间戳：%v,unixnano时间戳：%v\n", now.Unix(), now.UnixNano())

	// 示例
	// 统计一个程序运行了多长时间
	start := time.Now().Unix()
	fmt.Printf("start unix时间戳：%v\n", start)
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Printf("end unix时间戳：%v\n", end)
	fmt.Printf("一共花费了%v秒\n", end-start)

}
