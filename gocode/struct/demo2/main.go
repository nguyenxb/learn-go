package main

import "fmt"

type Point struct {
	x, y int
}
type Rect struct {
	leftUp, rightDown Point
}
type Rect2 struct {
	leftUp, rightDown *Point
}

type A struct {
	num  int
	num2 int
}
type B struct {
	num  int
	num2 int
}

// 两个结构体之间可以相互转化，但是结构体的字段个数，字段名称，
// 字段类型,定义的顺序，要完全一样才能转换
func main() {

	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// 结构体的存储空间在内存中是连续分布的
	fmt.Printf("r1.leftUp.x 地址=%p, r1.leftUp.y 地址=%p\n", &r1.leftUp.x, &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x 地址=%p, r1.rightDown.y 地址=%p\n", &r1.rightDown.x, &r1.rightDown.y)

	//  r2 的两个*Point类型是连续的，这两个*point本身的地址是连续的，
	// 但是他们所指向的地址不一定是连续的 ？？
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}
	fmt.Printf("r2.leftUp 本身地址=%p, r2.rightDown 本身地址=%p\n", &r2.leftUp, &r2.rightDown)
	fmt.Printf("r2.leftUp 指向的地址=%p, r2.rightDown 指向的地址=%p\n", r2.leftUp, r2.rightDown)
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)
}
