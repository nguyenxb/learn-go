package main

import (
	"errors"
	"fmt"
	"time"
)

func test() {
	// 使用defer+recover 来捕获和处理异常
	defer func() {
		// err := recover()                // recover()内置函数，可以捕获到异常
		if err := recover(); err != nil { // 说明捕获到错误
			fmt.Println("err=", err)
			// 在这里可以将错误信息发送给管理员。。
			fmt.Println("发送邮件给admin@sohu,cin")
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// 函数取读取配置文件init.conf的信息
// 如果文件名传入不正确，我们就返回一个自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误：..")
	}
}
func test02() {
	err := readConf("config.in")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02()继续执行")
}

func main() {
	test()
	for i := 1; i < 5; i++ {
		fmt.Println("main() 下的代码")
		time.Sleep(time.Second)
	}
	test02()
	fmt.Println("main() 。。。。。")

}
