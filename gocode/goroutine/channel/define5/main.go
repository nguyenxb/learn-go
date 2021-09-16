package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// 往管道中放任意类型的数据
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	// 存入
	cat1 := Cat{"tom", 14}
	cat2 := Cat{Name: "jack", Age: 56}
	allChan <- cat1
	allChan <- cat2
	allChan <- "小明"
	allChan <- 36

	// 取出
	cat11 := <-allChan
	cat22 := <-allChan
	v1 := <-allChan
	v2 := <-allChan

	//如果没有转成Cat类型 会报错: cat11.Name undefined
	// (type interface {} is interface with no methods)
	cat111, ok := cat11.(Cat) // 类型断言
	if ok {
		fmt.Println(cat111.Name)
	} else {
		fmt.Println("转换失败")
	}

	fmt.Println(cat11)
	fmt.Println(cat22)
	fmt.Println(v1)
	fmt.Println(v2)
}
