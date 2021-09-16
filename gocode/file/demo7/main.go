package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 统计文件中有多少个空格,英文，数字，其他字符
	srcFileName := "C:\\Users\\asd\\Desktop\\407.txt"
	//  打开文件
	srcFile, err := os.OpenFile(srcFileName, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer srcFile.Close()
	// 获取reader对象
	reader := bufio.NewReader(srcFile)
	var enCount int
	var NumCount int
	var spaceCount int
	var otherCount int
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				enCount++
			case v == ' ' || v == '\t':
				spaceCount++
			case v >= '0' && v <= '9':
				NumCount++
			default:
				otherCount++

			}
		}
	}
	fmt.Println("enCount=", enCount)
	fmt.Println("NumCount=", NumCount)
	fmt.Println("spaceCount=", spaceCount)
	fmt.Println("otherCount=", otherCount)

}
