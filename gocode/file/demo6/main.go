package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path1 := "C:\\Users\\asd\\Desktop\\JAVA综合设计实验2020.doc"
	path2 := "C:\\Users\\asd\\Desktop\\新建文件夹\\JAVA综合设计实验2020.doc"
	written, err := CopyFile(path2, path1)
	fmt.Println("written=", written)
	fmt.Println("err=", err)
}
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	// 打开文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer srcFile.Close()
	// 获取Reader
	reader := bufio.NewReader(srcFile)

	// 打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err =", err)
		return
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	fmt.Printf("srcFile type %T\n,srcFile =%v\n ", srcFile, srcFile)
	fmt.Printf("reader type %T \n,reader =%v \n", reader, reader)
	fmt.Printf("dstFile type %T\n,dstFile =%v \n", dstFile, dstFile)
	fmt.Printf("writer type %T\n,writer =%v \n", *writer, *writer)

	// 实现复制文件
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	writer.WriteString(str)
	// }
	return io.Copy(writer, reader)
}
