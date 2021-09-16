package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//  打开文件
	file, err := os.Open("C:\\Users\\asd\\Desktop\\407.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	// 查看文件的类型
	fmt.Printf("file = %v,file type = %T\n", file, file)

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件失败", err)
		}
	}()
	// 带缓冲区的读取
	reader := bufio.NewReader(file)

	for {
		// 表示读取到一个换行就结束本次读取
		str, err := reader.ReadString('\n')
		// 读取到文件末尾时退出
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("读取文件结束11。。")
	fmt.Println()

	// 直接读取整个文件，一次性直接读取完,并且将打开和关闭文件封装到
	// 函数中，只适合文件不大的时候
	fileStr := "C:\\Users\\asd\\Desktop\\407.txt"
	content, err1 := ioutil.ReadFile(fileStr)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("file = %s", content)

}
