package main

import "fmt"

func main() {
	// 切片：
	/*
		切片的底层：
		stuct {
			ptr *[5]int // 指针
			len int // 元素个数
			cap int // 容量
		}
	*/
	// 基于数组创建切片
	var arr = [...]int{23, 34, 45, 56, 65, 12, 42, 67, 98}
	var slice = arr[4:]
	slice1 := slice[3:5]
	fmt.Println("arr=\t", arr)
	fmt.Println("slice=\t", slice)
	fmt.Println("slice1=\t", slice1)
	fmt.Println("slice的元素个数：", len(slice))
	fmt.Println("slice的容量：", cap(slice))
	fmt.Println("slice1的元素个数：", len(slice1))
	fmt.Println("slice1的容量：", cap(slice1))

	arr[6] = 1000
	slice[3] = 700
	fmt.Println("arr", arr)
	fmt.Println("slice=", slice)

	var slice2 []int = make([]int, 3)
	fmt.Println("slice2=", slice2)

	//	 创建一个指向包含5个元素的空间，只能通过slice访问这个存储空间，对外是不可见的
	var slice3 []float64 = make([]float64, 5, 10)
	fmt.Println("slice3=", slice3)
	fmt.Println("slice3 len=", len(slice3))
	fmt.Println("slice3 cap=", cap(slice3))
	slice3[3] = 90
	var slice4 []float64 = slice3
	fmt.Println("slice4=", slice4)
	fmt.Printf("slice3 address=%p\n", &slice3)
	fmt.Printf("slice4 address=%p", &slice4)
	fmt.Println()
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("slice4[%v]=%v\n", i, slice4[i])
	}
	for ind, val := range slice4 {

		fmt.Printf("slice4[%v]=%v\n", ind, val)
	}

	slice3 = append(slice3, 100)
	slice3 = append(slice3, slice3...)
	slice3[1] = 1000
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Printf("slice3=%p", slice3)
	fmt.Printf("slice4=%p", slice4)
	fmt.Println()
	var slice5 []int = []int{1, 2, 3, 4, 5}
	var slice6 []int = make([]int, 10)
	copy(slice6, slice5)
	fmt.Println(slice5)
	fmt.Println(slice6)
}
