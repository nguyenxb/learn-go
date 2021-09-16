package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	// 类型断言
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	// b = a //不能直接赋值
	// 下面这条语句是类型断言，表示判断a 是否指向Point类型的
	// 变量，如果是则就转成Point类型并赋值给b,否则报错
	b = a.(Point) // 可以，这就是类型断言
	fmt.Println(b)

	var a1 interface{}
	var point1 Point = Point{3, 4}
	a1 = point1
	// 带检查的类型断言
	b1, ok := a1.(Point) // ok 是bool类型
	if ok {
		fmt.Println(b1)
	}
}
