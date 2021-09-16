package main

import (
	"fmt"
	"math"
)

// 向 intChan 放入1-8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
	// fmt.Println("存入管道完成")
}
func prime(intChan, primeChan chan int, exitChan chan bool) {
	var flag bool
	for v := range intChan {
		flag = true
		for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
			// fmt.Println("asd", i)
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag && v != 1 {
			primeChan <- v
		}
	}
	exitChan <- true
	//  第一种方法
	// if len(exitChan) == 4 {
	// 	close(exitChan)
	// 	close(primeChan)
	// }
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	// 标识退出的管道
	exitChan := make(chan bool, 4)
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go prime(intChan, primeChan, exitChan)
	}
	// 当且仅当4个协程全部完成时，才能退出

	// 第一种方法
	// for {
	// 	if len(exitChan) == 4 {
	// 		break
	// 	}
	// }

	//第二种方法
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan // 直接释放管道中的4个元素
		}
		close(primeChan)
	}()
	i := 1
	for v := range primeChan {
		fmt.Print(v, " ")
		if i%10 == 0 {
			fmt.Println()
		}
		i++
	}
	fmt.Println("完成i=", i)
}
