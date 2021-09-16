package main

import (
	"fmt"
)

type Pupil struct {
	Name  string
	Age   int
	Socre int
}

func (p *Pupil) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Socre)
}
func (p *Pupil) SetScore(score int) {
	p.Socre = score
}
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试。。")
}

type Graduate struct {
	Name  string
	Age   int
	Socre int
}

func (p *Graduate) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Socre)
}
func (p *Graduate) SetScore(score int) {
	p.Socre = score
}
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试。。")
}
func main() {
	// 继承 ， 解决代码冗余 详见 ../demo2
	var pupil = &Pupil{
		Name: "tom",
		Age:  10,
	}
	pupil.testing()
	pupil.SetScore(69)
	pupil.ShowInfo()
	var graduate = &Graduate{
		Name: "tom",
		Age:  10,
	}
	graduate.testing()
	graduate.SetScore(69)
	graduate.ShowInfo()
}
