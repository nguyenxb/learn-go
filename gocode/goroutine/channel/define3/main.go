package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//  管道中存放结构体类型
	var catChan chan Cat
	// 该管道最多可以放10个Cat 类型的元素
	catChan = make(chan Cat, 10)

	cat1 := Cat{Name: "tom1", Age: 15}
	cat2 := Cat{Name: "jack", Age: 28}

	// 往管道中放入cat元素
	catChan <- cat1
	catChan <- cat2

	// 取出
	cat11 := <-catChan
	cat22 := <-catChan

	fmt.Println(cat11)
	fmt.Println(cat22)
}
