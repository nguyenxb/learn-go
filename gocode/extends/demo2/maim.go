package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Socre int
}

func (p *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Socre)
}
func (p *Student) SetScore(score int) {
	p.Socre = score
}

// 给*Student增加一个方法
func (p *Student) GetSum(num1, num2 int) int {
	return num1 + num2
}

type Graduate struct {
	Student //嵌入Student匿名结构体
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试。。")
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试..")
}

func main() {
	// 继承特性
	var graduate = &Graduate{}
	graduate.Student.Name = "tom"
	graduate.Student.Age = 18
	graduate.Student.Socre = 66
	graduate.testing()
	graduate.Student.SetScore(89)
	graduate.Student.ShowInfo()
	fmt.Println("大学生：", graduate.Student.GetSum(10, 60))

	pupil := &Pupil{}
	pupil.Student.Name = "小红"
	pupil.Student.Age = 9
	pupil.testing()
	pupil.Student.SetScore(99)
	pupil.Student.ShowInfo()
	fmt.Println("小学生", pupil.Student.GetSum(10, 15))

}
