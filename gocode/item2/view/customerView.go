package main

import (
	"fmt"
	"item2/model"
	"item2/service"
)

type customerView struct {
	// 用户输入
	key string
	// 记录是否退出系统
	loop bool
	// 绑定一个处理文件
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("-----显 示 客 户 列 表-----")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, value := range customers {
		fmt.Println(value.Info())
	}
	fmt.Println("-------显 示 完 成---4----\n\n")
}
func (this *customerView) add() {
	fmt.Println("添加客户信息:")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}

}
func (this *customerView) delete() {
	fmt.Println("删除客户信息")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	chioce := ""
	for {
		fmt.Println("是否确认删除（y/n）：")
		fmt.Scanln(&chioce)
		if chioce == "y" || chioce == "Y" || chioce == "N" || chioce == "n" {
			break
		}
		fmt.Println("输入错误")
	}
	if chioce == "y" || chioce == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败")
		}
	}
}
func (this *customerView) upDate() {
	fmt.Println("修改客户信息")
	fmt.Println("请选择待修改客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer1(id, name, gender, age, phone, email)
	chioce := ""
	for {
		fmt.Println("是否确认修改（y/n）：")
		fmt.Scanln(&chioce)
		if chioce == "y" || chioce == "Y" || chioce == "N" || chioce == "n" {
			break
		}
		fmt.Println("输入错误")
	}
	if chioce == "y" || chioce == "Y" {
		if this.customerService.Update(customer) {
			fmt.Println("修改成功")
		} else {
			fmt.Println("修改失败")
		}
	}
}

func (this *customerView) MainMenu() {

	for {
		fmt.Println("-----客 户 管 理 系 统------")
		fmt.Println("     1. 添加客户信息")
		fmt.Println("     2. 修改客户信息")
		fmt.Println("     3. 删除客户信息")
		fmt.Println("     4. 显示客户列表")
		fmt.Println("     5. 退出系统")
		fmt.Print("请选择操作（1-5）：")
		fmt.Scanln(&this.key)
		fmt.Println()
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.list()
			this.upDate()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			chioce := ""
			for {
				fmt.Println("是否确认退出系统（y/n）:")
				fmt.Scanln(&chioce)
				if chioce == "y" || chioce == "y" ||
					chioce == "N" || chioce == "n" {
					break
				}
			}
			if chioce == "Y" || chioce == "y" {
				this.loop = false
			}
		default:
			fmt.Println("无此选项")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您已退出客户管理系统")
}

func main() {
	// 创建一个界面对象
	customerView := customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}
	// customerView.customerService = service.NewCustomerService()
	customerView.MainMenu()
	// arr := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(find(arr, 3))
}

// func find(arr []int, id int) int {
// 	min := 0
// 	max := len(arr) - 1
// 	mid := 0
// 	for min <= max {
// 		mid = min + (max-min)>>1
// 		if arr[mid] > id {
// 			fmt.Println(mid)
// 			max = mid - 1
// 		} else if arr[mid] < id {
// 			min = mid + 1
// 		} else {
// 			return mid
// 		}
// 	}
// 	return -1 - mid
// }
