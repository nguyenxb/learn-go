package main

import "fmt"

// 声明，定义接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}
type Usb2 interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
}

// 让Phone 实现 Usb 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

// 接收一个Usb 接口类型的变量
//  所谓实现Usb接口，就是实现了Usb接口声明的所有方法
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

// 比如上面的Phone Camera就是实现了接口 Usb 和 Usb2 两个接口
func main() {
	// 多态是使用接口来实现的
	// 先创建对象
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//多态的体现
	computer.Working(phone)
	computer.Working(camera)
}
