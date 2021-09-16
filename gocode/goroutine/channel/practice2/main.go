package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

func writeDataToFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("writeDataToFile 文件打开失败")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := ""
	for i := 0; i < 1000; i++ {
		str = strconv.Itoa(rand.Intn(1000) + 1)
		_, err = writer.WriteString(str + " ")
		if err == io.EOF {
			break
		}
	}
	fmt.Println("文件写入完成")
}
func sort(fileName string, numChan chan int, exitChan chan bool) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("sort 文件打开失败")
		return
	}
	defer file.Close()
	numStrChan := make(chan string, 1000)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		numStrChan <- str
		if err != io.EOF {
			break
		}
	}
	close(numStrChan)
	fmt.Println("sort 文件读取完成")

	for {
		numStr := <-numStrChan
		num, _ := strconv.Atoi(numStr)
		numChan <- num
		if len(numStrChan) == 0 {
			break
		}
	}
	fmt.Println("numChan 管道加载完成")
	close(numChan)
	numArr := make([]int, len(numChan))
	var i int
	writer := bufio.NewWriter(file)
	for v := range numChan {
		numArr[i] = v
		i++
	}
	QuickSort(numArr)
	for i := 0; i < len(numArr); i++ {
		writer.WriteString(strconv.Itoa(numArr[i]) + " ")
	}
	exitChan <- true
	close(exitChan)
}
func quickSort(arr []int, left, right int) {
	// 中心轴
	pivot := arr[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && arr[j] >= pivot {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}
		if i <= p && arr[i] <= pivot {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = pivot
	if p-left > 1 {
		quickSort(arr, left, p-1)
	}
	if right-p > 1 {
		quickSort(arr, p+1, right)
	}
}
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
func main() {
	/* 1,开一个协程，writeDataToFile,随机生成1000 个数据，存放到文件中
	2 当writeDataToFile 完成 写 1000 个数据到文件后，让sort协程从文件中读取
	1000 个数据，并完成排序，重新写入到另一个结果文件

	*/
	numChan := make(chan int, 1000)
	exitChan := make(chan bool, 1)
	fileDir := "C:\\Users\\asd\\Desktop\\test\\"
	// fileDir1 := "C:\\Users\\asd\\Desktop\\test1\\"
	writeDataToFile(fileDir + "num.txt")
	sort(fileDir+"num.txt", numChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
	fmt.Println("完成")
}
