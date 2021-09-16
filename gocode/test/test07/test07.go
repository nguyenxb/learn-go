package main

import "fmt"

func main() {
	var (
		height int16
		weight int16
	)
	fmt.Println("请输入身高：")
	fmt.Scanln(&height)
	fmt.Println("请输入体重：")
	fmt.Scanln(&weight)

	if weight-10 <= (height-108)*2 && (height-108)*2 >= weight {
		fmt.Println("合适")
	} else {
		fmt.Println("不合适")
	}
}
