package main

import (
	_ "bufio"
	"fmt"
	"os"
)

func main() {
	// 通过os.OpenFile()函数打开文件
	// 打开文件 fileStr
	fileStr := "C:\\Users\\asd\\Desktop\\407.txt"
	file, err := os.OpenFile(fileStr, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=%v\n")
	}
	defer file.Close()
	// 如果文件中有数据，会直接覆盖原来的数据
	// 往文件中写入5句话，Hello, Gardon
	str := "Hello, Gardonaaaaa"
	// 从文件头开始写入
	ret, err1 := file.WriteString(str)
	if err1 != nil {
		fmt.Println("err1", err1)
	}
	fmt.Println(len(str))
	fmt.Println(ret)

	// 第二种方式
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str + "asd\r\n")
	}
	// 将缓冲区的数据刷新到文件中
	writer.Flush()
	fmt.Println("writer")
}
