package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用select 可以解决从管道取数据的阻塞问题

	// 1. 定义一个管道 int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 2.管道 string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("hello,%v", i)
	}
	// 传统的方法在遍历管道时，如果不关闭阻塞会导致deadlock

	// 问题是，我们在实际开发中，可能不号确定什么时候关闭该
	// 通道，因此我们可以使用select 方式解决

	for {
		select {
		case v := <-intChan:
			time.Sleep(time.Second)
			fmt.Printf("从intChan读取的数据%v\n", v)
		case v := <-stringChan:
			time.Sleep(time.Second)
			fmt.Printf("从stringChan读取的数据%v\n", v)
		default:
			time.Sleep(time.Second)
			fmt.Println("两者都读取不到数据了")
			// break //  跳出select.
			return
		}
		// break // 跳出 for
	}
}
