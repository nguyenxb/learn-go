package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 10
	intChan <- 200
	close(intChan)
	// 关闭管道后只能取，不能存
	// panic: send on closed channel
	// intChan <- 33 // error

	num1 := <-intChan
	num2 := <-intChan

	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	intChan2 := make(chan int, 100)
	// 往管道中放入100个数据
	for i := 0; i < 100; i++ {
		intChan2 <- i
	}
	// len(intChan2)返回的数是不确定的，使用for遍历不一定能
	// 取完其中的全部元素,
	// 使用这种方法遍历实际上管道有 100个元素，只能取50个，不正确
	// for i := 0; i < len(intChan2); i++ {
	// 	num := <-intChan2
	// 	fmt.Println("num[", i, "]", num)
	// }

	// 该方法可以全部取完元素
	//  对管道使用for -range 必须要关闭管道，
	// 否则会报错 ：fatal error: all goroutines are asleep - deadlock!
	// 因为关闭管道后，管道会有一个标志，没有关闭管道时，for - range 会一直取管道
	// 中的值，即fatal error: all goroutines are asleep - deadlock!
	// 关闭管道 func close(c chan<- Type)
	close(intChan2)

	// 对管道进行for-range 遍历，只会返回一个值，值类型就是管道定义的类型
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
	fmt.Println("取出管道intChan2的全部元素")
}
