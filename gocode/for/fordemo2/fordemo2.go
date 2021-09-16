package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i < 8; i++ {
		for j := i; j < 8; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == 7 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i := 1; i < 5; i++ {
		for j := 1; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i < 8; i++ {
		if i <= 4 {
			for j := i; j < 4; j++ {
				fmt.Print(" ")
			}
			for k := 1; k <= 2*i-1; k++ {
				if k == 1 || k == 2*i-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}

		} else {
			for j := 1; j <= i-4; j++ {
				fmt.Print(" ")
			}
			for k := 1; k <= 8-2*(i-4)-1; k++ {
				if k == 1 || k == 8-2*(i-4)-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}

		}
		fmt.Println()
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 1; i <= 9; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("\t\t")
		}
		for j := i; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

}
