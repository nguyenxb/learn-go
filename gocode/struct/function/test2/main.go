package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

type Box struct {
	length float64
	width  float64
	height float64
}

func (box *Box) getVolum() float64 {
	return box.length * box.width * box.height
}

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) getPrice() {
	if visitor.Age > 18 {
		fmt.Println("门票价格：", 20)
	} else {
		fmt.Println("免费")
	}
}

func main() {
	//1
	student := Student{"小明", "男", 18, 1020, 88}
	fmt.Println(student.say())

	// 2
	box := Box{3.2, 5.6, 8}
	fmt.Printf("体积=%.2f", box.getVolum())
	fmt.Println()
	// 3
	visitor2 := Visitor{
		Age:  19,
		Name: "asd",
	}
	visitor := Visitor{"小明", 20}
	visitor1 := Visitor{"小明", 8}
	visitor.getPrice()
	visitor1.getPrice()
	visitor2.getPrice()
}
