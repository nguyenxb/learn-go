package main

import "fmt"

func main() {
	/* 1,开 10 个协程，writeDataToFile,每个协程随机生成1000 个数据，存放到 10个 文件中
	2 当writeDataToFile 完成 写 1000 个数据到 10 个文件后，让 10个 sort协程从10个文件中读取
	1000 个数据，并完成排序，重新写入到另10 个结果文件

	*/
	numChan := make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		numChan <- i
	}
	close(numChan)
	for v := range numChan {
		fmt.Println(v, "::", len(numChan))
	}
	fmt.Println(len(numChan))
}
