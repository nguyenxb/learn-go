package main

import "fmt"

func main() {
	var score uint8
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)
	switch {
	case 90 <= score && score <= 100:
		fmt.Printf("%d 成绩优秀", score)
	case 80 <= score:
		fmt.Printf("%d 成绩优良", score)
	case 70 <= score:
		fmt.Printf("%d 成绩良好", score)
	case 60 <= score:
		fmt.Printf("%d 成绩合格", score)
	default:
		fmt.Printf("%d 成绩不合格", score)
	}

}
