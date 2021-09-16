package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：计算1-200的各个数的阶乘，并且把各个数的阶乘放入map中
// 最后显示出来，要求使用goroutine完成
var (
	myMap = make(map[int]uint, 10)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁
	// sync 是包： synchronized 同步
	// Mutex ：是互斥
	lock sync.Mutex
)

// test 函数就是计算n!,将整个结果放入到 myMap 中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = uint(res) //问题1： error :concurrent map writes
	lock.Unlock()
}

//	在编译时 使用 -race 可以查看是否有资源竞争
// go build -race main.go
// main.exe
func main() {
	// 开启200个协程来完成这个任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	// 休眠10秒，等待协程运行完
	// 问题2： 不知道该等多久
	// 休眠时间长了，会加长等待时间，短了可能goroutine还在执行，
	// 但主线程结束了，协程也跟着结束了，造成程序没有执行完就结束了
	time.Sleep(5 * time.Second)

	// 为什么要在这里加锁？ 为了避免资源竞争

	lock.Lock()
	// 输出map中的结果
	for i, v := range myMap {
		fmt.Println("map[", i, "]=", v)
	}
	lock.Unlock()
}
