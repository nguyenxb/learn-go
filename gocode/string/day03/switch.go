// switch
package main

import (
	"fmt"
)

func main() {
	var num = 3
	// compile error
	// switch num {
	// case num > 5:
	// 	fmt.Println("aaa")
	// case num > 3:
	// 	fmt.Println("nnn")
	// default:
	// 	fmt.Println("mmm")
	// }

	for i, num := 0, 0; num <= 11; i, num = i+1, num+1 {

		// switch i {
		// case 0:
		// 	fmt.Println("i =", i, "output : 0")
		// case 1:
		// 	fmt.Println("i =", i, "output : 1")
		// case 2:
		// 	fallthrough
		// case 3:
		// 	fmt.Println("i =", i, "output : 3")
		// case 4, 5, 6:
		// 	fmt.Println("i =", i, "output : 4,5,6")
		// default:
		// 	fmt.Println("i =", i, "output : default")
		// }
		switch true {
		case 0 <= num && num <= 3:
			fmt.Println("num =", num, "output: 0-3")
		case 4 <= num && num <= 6:
			fmt.Println("num =", num, "output: 4-6")
		case 7 <= num && num <= 9:
			fmt.Println("num =", num, "output: 7-9")
		default:
			fmt.Println("num =", num, "output: greater than or equal to 10 ")
		}
	}

	sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println("sum =", sum)

Jloop:
	// the break statement terminates(终止) the other outer（外面的） loop at the Jloop tag（标签）
	for j := 0; j < 5; j++ {
	inloop:
		for i := 0; i < 10; i++ {
			if i > 4 {
				break Jloop
			}
			fmt.Println(i)
			break inloop
			fmt.Println(i)
		}
	}
	myfunc()

}
func myfunc() {
	i := 0
Here:
	fmt.Println("myfunc : i =", i)
	i++
	if i < 10 {
		goto Here
	}
}
