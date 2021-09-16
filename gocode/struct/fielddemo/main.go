package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	score [5]float64
	ptr   *int
	slice []int
	map1  map[string]string
}

func main() {
	var p1 Person
	fmt.Println("p1 = ", p1)

	if p1.ptr == nil {
		fmt.Println("p1.ptr = ", p1.ptr, " ok1")
	}
	if p1.slice == nil {
		fmt.Println("p1.slice=", p1.slice)
	}
	if p1.map1 == nil {
		fmt.Println("p1.map1=", p1.map1)
	}
	// error : index out of range [0] with length 0
	// p1.slice[0] = 1 // 没有创建空间

	// 使用 map ,数组切片必须先创建空间
	p1.slice = make([]int, 10)
	p1.slice[0] = 11
	fmt.Println("p1.slice 2 = ", p1.slice)

	// 必须要创建空间
	p1.map1 = make(map[string]string, 2)
	p1.map1["123"] = "asd"
	p1.map1["111"] = "asdad"
	fmt.Println("map1 2 = ", p1.map1)

	// 直接初始化
	// var i int
	var p2 = Person{"na", 11, [5]float64{1, 2, 3, 4, 5}, new(int),
		make([]int, 4, 9), make(map[string]string, 3)}
	fmt.Println("p2=", p2)

	var p3 *Person = new(Person)

	/// 底层 调用的是 (*p3).Age
	// 也等价于 p3.Age ，
	// 这种写法在底层会给 p3.Age 加上取值运算 (*p3).Name
	//  简化写法
	fmt.Println(p3.Age)
	// 标准写法
	fmt.Println((*p3).Age)

	// 也可以这样定义,原因同上
	var p4 *Person = &Person{} // 可以在括号中赋值
	fmt.Println(p4)

	// 也可以
	var p5 *Person = &Person{"na", 11, [5]float64{1, 2, 3, 4, 5}, new(int),
		make([]int, 4, 9), make(map[string]string, 3)}
	fmt.Println(*p5)
	fmt.Println(p5)
}
