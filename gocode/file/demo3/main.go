package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//  清空文件再写
	fileStr := "C:\\Users\\asd\\Desktop\\407.txt"
	file, err := os.OpenFile(fileStr, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err=%v\n")
	}
	defer file.Close()
	str := "你好，，，\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
