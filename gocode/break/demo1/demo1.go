package main

import (
	"math/rand"
	"time"

	"fmt"
)

func main() {

	// time.Now().Unix(): 返回一个从1970年1月1日到0时0分0秒到现在的秒数
	// rand.Seed(time.Now().Unix())
	// 伪随机数
	// rand.Intn(100) + 1 // 随机生成 1 到100 的整数
	var count int
	for {
		// 伪随机数
		rand.Seed(time.Now().UnixNano()) //纳秒数
		time.Sleep(2)
		num := rand.Intn(100) + 1 // 随机生成 1 到100 的整数
		fmt.Println("num=", num)
		count++
		if num == 99 {
			fmt.Println("count=", count)
			break
		}
	}
}
