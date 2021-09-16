package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var flag bool = true

func writeData(numChan chan int) {
	for i := 0; i < 20000; i++ {
		numChan <- i
		fmt.Println(" writeData 写入数据：", i)
	}
	// 关管道
	close(numChan)
}
func readData(numChan, resChan chan int, exitChan chan bool, count int) {
	for v := range numChan {

		sum := 0
		for i := 0; i <= v; i++ {
			sum += v
		}
		resChan <- sum
		fmt.Println("readData", count, "sum=", sum)
		fmt.Println("readData", count, " 读取数据 ：", v)
	}
	// lock.Lock()
	fmt.Println("协程", count, "进来")

	// 如果没有flag标记会导致异常
	// 当最后一个协程读取完numChan 中的数据后，关闭了exitChan,而其他的协程进来时
	// 如果再次往exitChan再次写入数据会导致异常
	if flag {
		flag = false
		exitChan <- true
		close(exitChan)
	}
	fmt.Println("协程", count, "退出")
	// lock.Unlock()
}

func main() {
	/* 1 启动一个协程，将1-20000的数放入到一个channel中，如numChan
	   2 启动8个协程，从numChan中取出数据，并计算1 + 。。。n的值，
	并且存放到resChan中
		3 最后8个协程协调工作完后，在遍历resChan，显示结果[如res[1]=1,..res[20]=44]
	*/

	//  创建一个管道
	numChan := make(chan int, 20000)
	resChan := make(chan int, 20000)
	exitChan := make(chan bool, 1)
	// 创建一个 writeData 协程
	go writeData(numChan)
	// 创建8个readData协程
	for i := 0; i < 8; i++ {
		go readData(numChan, resChan, exitChan, i+1)
	}
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
	close(resChan)
	i := 1
	for v := range resChan {
		if i <= 20 {
			if v == 55 {
				fmt.Println(10, ".......................")
			}
			time.Sleep(time.Second)
		}
		fmt.Println("res[", i, "]=", v)
		i++
	}

	if len(numChan) == 0 {
		fmt.Println("执行完成", len(numChan))
	} else {
		fmt.Println("执行失败", len(numChan))
	}

}
