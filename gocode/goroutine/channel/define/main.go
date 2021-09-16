package main

import (
	"fmt"
)

func main() {
	// 演示管道的使用
	// 1, 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 2.intChan的值和类型
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	fmt.Printf("intChan 的类型=%T\n", intChan)

	// 3 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num // 把 num 存到管道中
	intChan <- 50
	// fmt.Println("intChan=", intChan)

	// 管道的长度和容量 (管道的容量是不能自动增长的)
	fmt.Printf("channel len=%v cap=%v", len(intChan), cap(intChan))

	// 从管道中读取数据
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println("num2=", num2, "num3=", num3)
	fmt.Printf("channel len=%v cap=%v", len(intChan), cap(intChan))
}
