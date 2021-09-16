package main

import (
	"fmt"
	"time"
)

//  如果写的速度慢而读取的速度快，那么管道会自动阻塞，直到有数据进来
// 管道是引用类型的
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		time.Sleep(500 * time.Millisecond)
		// 往管道中写入数据
		intChan <- i
		fmt.Println("writeData run ", i)
	}
	// 写完数据了直接关闭管道
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for v := range intChan {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("readData 读数据 ", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)

	// 如果不加这个for循环，协程没有执行结束时 ，主线程就会直接执行结束
	for {
		// 当无法从exitChan中没有数据时，会返回ok = false
		// 直到 readData 函数往管道中写入数据时，主线程才结束
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
	fmt.Println("执行结束")
}
