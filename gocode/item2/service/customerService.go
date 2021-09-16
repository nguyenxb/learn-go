package service

import (
	"item2/model"
)

type CustomerService struct {
	//  存储客户的编号
	customerNum int
	// 存储客户的信息
	customers []model.Customer
}

//  返回指针实例 *CustomerService
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customer := model.NewCustomer("张三", "男", 25, "123", "zhangsan@qq.com")
	customerService.customerNum = 1
	customer.Id = customerService.customerNum
	customerService.customers = append(customerService.customers, customer)
	return customerService
}
func (this *CustomerService) List() []model.Customer {
	return this.customers
}
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Delete(id int) bool {
	index := this.findById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
func (this *CustomerService) findById(id int) int {
	min := 0
	max := len(this.customers) - 1
	mid := 0
	for min <= max {
		mid = min + (max-min)>>1
		if this.customers[mid].Id > id {
			max = mid - 1
		} else if this.customers[mid].Id < id {
			min = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func (this *CustomerService) Update(customer model.Customer) bool {
	index := this.findById(customer.Id)
	if index == -1 {
		return false
	}
	if customer.Name != "" {
		this.customers[index].Name = customer.Name
	}
	if customer.Gender != "" {
		this.customers[index].Gender = customer.Gender
	}
	if customer.Age != 0 {
		this.customers[index].Age = customer.Age
	}
	if customer.Phone != "" {
		this.customers[index].Phone = customer.Phone
	}
	if customer.Email != "" {
		this.customers[index].Email = customer.Email
	}
	return true
}
