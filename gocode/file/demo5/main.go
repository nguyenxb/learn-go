package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 经一个文件的数据 复制 到另一个文件中
	fileStr := "C:\\Users\\asd\\Desktop\\407.txt"
	fileStr2 := "C:\\Users\\asd\\Desktop\\999.txt"

	// 直接读取整个文件
	data, err := ioutil.ReadFile(fileStr)
	if err != nil {
		fmt.Println("read file err=%v", err)
		return
	}

	err = ioutil.WriteFile(fileStr2, data, 0777)
	if err != nil {
		fmt.Println("write file error = %v", err)
	}

	b1, err1 := PathExists(fileStr)
	b2, err2 := PathExists(fileStr2)
	b3, err3 := PathExists(fileStr + "dadsasd")
	fmt.Println("b1=", b1, "err=", err1)
	fmt.Println("b1=", b2, "err=", err2)
	fmt.Println("b1=", b3, "err=", err3)

}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
