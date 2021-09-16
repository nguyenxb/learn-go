package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取本机的cpu数量
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	// 设置该程序使用的cpu数目
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
