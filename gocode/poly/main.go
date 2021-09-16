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
	fmt.Println(this.Name, "phone start run")
}
func (this Phone) Stop() {
	fmt.Println(this.Name, "Phone stop ")
}

type Camera struct {
	Name string
}

func (this Camera) Start() {
	fmt.Println(this.Name, "camera start run")
}
func (this Camera) Stop() {
	fmt.Println(this.Name, "camera stop ")
}

func main() {
	// 多态数组
	var usbs [3]Usb
	fmt.Println(usbs)
	usbs[0] = Phone{"小米"}
	usbs[1] = Phone{"华为"}
	usbs[2] = Camera{"索尼"}
	for i, _ := range usbs {
		usbs[i].Start()
		usbs[i].Stop()
	}
	fmt.Println(usbs)
}
