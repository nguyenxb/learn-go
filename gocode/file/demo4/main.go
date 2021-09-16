package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//  以追加的方式写入文件
	fileStr := "C:\\Users\\asd\\Desktop\\407.txt"
	// 要更改模式
	file, err := os.OpenFile(fileStr, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err=%v\n")
	}
	defer file.Close()
	// 先读取文件
	reader := bufio.NewReader(file)
	for {
		str1, err3 := reader.ReadString('\n')
		if err3 == io.EOF {
			break
		}
		fmt.Print(str1)
	}

	str := "你好aaaaa\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
