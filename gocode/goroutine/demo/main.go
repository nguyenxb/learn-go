package main

import (
	"fmt"
	"strconv"
	"time"
)

type monster struct {
	num int
}

func test(this *monster) {
	for i := 1; i < 10; i++ {
		this.num++
		fmt.Println("test() hello,world.." + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	fmt.Println(this.num)
}

func main() {
	// go test()
	monster := &monster{0}
	go test(monster)
	for i := 1; i < 10; i++ {
		fmt.Println("main() run.." + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	for i := 0; i < 10; i++ {
		monster.num += 2
	}

	fmt.Println(monster.num)
}
