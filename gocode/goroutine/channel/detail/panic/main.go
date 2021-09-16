package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("sayHello,", i)
	}
}
func test() {
	// 处理异常后，即使该协程出现异常，也不会造成整个程序停止
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error,:", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "gitttt"
}

//没有任何处理的时候,当一个协程出现异常的时候,整个程序会直接结束
func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main(),", i)
		time.Sleep(time.Second)
	}

}
