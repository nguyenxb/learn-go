package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("参数个数", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("第%v个参数：%v\n", i, v)
	}
}
