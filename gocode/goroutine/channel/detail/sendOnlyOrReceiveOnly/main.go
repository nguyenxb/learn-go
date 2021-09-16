package main

import "fmt"

func main() {
	// 双向管道，可读可写
	var chan1 chan int
	chan1 = make(chan int, 10)
	chan1 <- 10
	chan1 <- 99
	chan1 <- 10
	chan1 <- 99
	chan1 <- 10

	num1 := <-chan1
	num2 := <-chan1
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	// 单向管道。只能读
	var chan2 <-chan int
	chan2 = chan1 // 获取chan1的地址
	num1 = <-chan2
	num2 = <-chan2
	// error:(send to receive-only type <-chan int)
	// chan2 <- 100 // error
	// num3 = <-chan2
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	// 单向管道， 只能写
	var chan3 chan<- int
	// 对于只读只写，只是管道的属性，不是管道的类型，
	// 在使用make创建内存空间的时候还是使用make(chan int,10)
	chan3 = make(chan int, 6)
	chan3 <- 10
	chan3 <- 1
	chan3 <- 13
	chan3 <- 14
	// <-chan3 (receive from send-only type chan<- int)
	// num1 = <-chan3 // error
	// num2 = <-chan3 // error
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	//  应用实例
	fmt.Println()
	var chan4 chan int
	chan4 = make(chan int, 8)
	writeOnly(chan4)
	readOnly(chan4)
}

// 该函数传入的管道只支持写入
func writeOnly(chan4 chan<- int) {
	chan4 <- 999
	// num := <-chan4 // error
	fmt.Println("writeOnly")
}

// 该函数传入的管道只支持读取
func readOnly(chan4 <-chan int) {
	// chan4 <- 100 // error
	num := <-chan4
	fmt.Println("readOnly", num)
}
