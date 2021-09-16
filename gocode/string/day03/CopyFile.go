package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
func main() {
	w, err := CopyFile("C:\\Users\\asd\\Desktop\\2.txt", "C:\\Users\\asd\\Desktop\\1.txt")
	fmt.Println(w, ":", err)
}
