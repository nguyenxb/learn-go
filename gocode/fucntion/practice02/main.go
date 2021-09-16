package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成一个随机数
func guess() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(99) + 1
}

func main() {
	var num int = guess()
	var input int
	fmt.Println(num)
myfor:
	for i := 1; i <= 10; i++ {
		fmt.Println("请输入一个1-100的整数:")
		fmt.Scanln(&input)
		if input == num {
			switch i {
			case 1:
				fmt.Println("猜对了,你真是天才")
				break myfor
			case 2, 3:
				fmt.Println("猜对了,你真聪明")
				break myfor
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("猜对了,一般般")
				break myfor
			case 10:
				fmt.Println("可算猜对了")
				break myfor
			}
		} else if input > num {
			fmt.Println("猜大了")
		} else if input < num {
			fmt.Println("猜小了")
		} else {
			fmt.Println("没有猜数机会了")
		}
	}
}
