package main

import "fmt"

// 编译通过
var Age int = 20

//  报错
// Name:= "tom" //compile error, 等价于 var Name string; Name = "tom"

func sum(n1, n2 int) int {
	// 当执行defer语句时，暂时不执行，会将defer后面的语句压入
	// 到独立的栈中，同时将当前值拷贝，当函数执行完毕后，再从defer栈
	// 按照先入后出的方式出栈执行
	defer fmt.Println("ok1 n1=", n1) // 3
	defer fmt.Println("ok2 n2=", n2) // 2
	n1++                             // n1 = 11
	n2++                             // n2 = 21
	res := n1 + n2
	fmt.Println("ok3 res =", res) // 1
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("main() res=", res) // 4
	//	作用域
	{
		res := 1
		fmt.Println(res)
	}
	/*
		使用示例（模拟代码）：
		// 关闭文件资源
		file = openfile(文件名)
		defer file.close()
		// 其他代码


		// 释放是数据库资源
		connect = openDatabse()
		defer connect.close()
		// 其他代码

	*/
}
