package main

import "fmt"

func main() {
	// string 底层是一个byte数组，
	//因此string也可以进行切片处理
	str := "hello@atguigu"
	// var str1 []byte = str[:5] // 编译不通过
	var str1 string = str[:5] // 编译通过，可简化如下
	slice := str[6:]
	fmt.Println("str", str)
	fmt.Printf("slice=%v,type=%T\n", slice, slice)
	fmt.Println("str1=", str1)

	// 要修改字符串，可以先把string 转化为 []byte 或者 []rune
	// 进行修改，然后重新转成 string
	var byteStr []byte = []byte(str)
	byteStr[0] = 'M'
	var str2 = string(byteStr)
	str = string(byteStr)
	fmt.Println("str1=", str2)
	fmt.Println("str=", str)

	// 细节 我们转化成[]byte后,可以处理英文和数字，但是不能处理中文
	// 原因是 []byte 是按字节来处理的，而一个汉字是三个字节，因此会出现
	// 乱码，解决方法是将 string 转成 []rune 即可，因为[]rune 是按字符来处理的，兼容汉字
	arr1 := []rune(str)
	arr1[0] = '被'
	str = string(arr1)
	fmt.Println("str=", str)
}
