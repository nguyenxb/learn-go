package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

func (this Phone) Start() {
	fmt.Println(this.Name, "手机开始工作")
}
func (this Phone) Stop() {
	fmt.Println(this.Name, "手机停止工作")
}
func (this Phone) Call() {
	fmt.Println(this.Name, "手机打电话")
}

type Camera struct {
	Name string
}

func (this Camera) Start() {
	fmt.Println(this.Name, "相机开始工作")
}
func (this Camera) Stop() {
	fmt.Println(this.Name, "相机停止工作")
}

type Computer struct {
}

func (this Computer) Working(usb Usb) {
	usb.Start()
	// 如果usb 是指向Phone 结构体变量，则还需要调用call方法
	// 类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var con Computer
	var phone1 = Phone{"小米"}
	var phone2 = Phone{"华为"}
	var camera1 = Camera{"索尼"}
	con.Working(phone1)
	con.Working(phone2)
	con.Working(camera1)

}
