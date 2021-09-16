package main

import "fmt"

func main() {
	for num, sum := 1, 0; ; num++ {
		sum += num
		if sum > 20 {
			fmt.Println("sum=", sum, "num=", num)
			break
		}
	}

	var (
		name     string
		password string
	)
	for i := 0; i < 3; i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		if name == "张三" && password == "1234" {
			fmt.Println("登入成功")
			break
		} else {
			fmt.Println("您已输入", i+1, "次错误, 最多只能输入3次")
		}
	}

here:
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if i == 2 {
				continue here
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}
