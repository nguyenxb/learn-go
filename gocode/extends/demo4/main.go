package main

import "fmt"

type A struct {
	Name string
	age  int
}
type B struct {
	Name  string
	Score float64
}
type C struct {
	A
	B
}

// 这就是组合
type D struct {
	a A
}

/*
	如果结构体中嵌套的是有名结构体，必须要带上有名结构体的名字
	这种关系是组合关系
*/
func main() {
	/*
		如果c 没有字段，而A和B有Name 有字段，
		这时必须要通过匿名结构体名字来区分
		所以c.Name 就会编译报错，方法也是一样
	*/
	var c C
	// c.Name = "tom"//.\main.go:20:3: ambiguous selector c.Name
	c.A.Name = "jack"
	c.B.Name = "mary"
	fmt.Println(c)

	var d D
	d.a.Name = "lili" // 有名结构体必须带上结构体的名字
	fmt.Println(d)
}
