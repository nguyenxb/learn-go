package main

import "fmt"

func main() {
	for i := 0; i < 26; i++ {
		if i%5 == 0 {
			fmt.Println()
		}
		fmt.Print(string('a'+i), " ")
	}
	for i := 0; i < 26; i++ {
		if i%5 == 0 {
			fmt.Println()
		}
		fmt.Print(string('Z'-i), " ")
	}
}
