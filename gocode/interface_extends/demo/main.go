package main

import (
	"fmt"
)

type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

type Monkey struct {
	Name string
}

//  继承
type littleMonkey struct {
	Monkey
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}
func (this *Monkey) Flying() {
	fmt.Println(this.Name, "通过学习 学会飞")
}
func (this *Monkey) Swimming() {
	fmt.Println(this.Name, "通过学习，学会游泳")
}

/*
	实现接口是对继承的补充
*/
func main() {
	var monkey = &Monkey{
		Name: "悟空",
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()

	//悟空变成鸟飞
	var bird BirdAble = monkey
	bird.Flying()

	// 悟空变成鱼游泳
	var fish FishAble = monkey
	fish.Swimming()

}
