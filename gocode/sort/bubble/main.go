package main

import "fmt"

func bubble(arr *[8]int) {
	flag := true

	for i := 0; i < len(arr)-1; i++ {
		flag = true

		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}

/*
	快速排序算法思想：
	1，选取一个中心轴pivot
	2,将大于pivot的数字放在pivot的右边
	3,将小于pivot的数字放在pivot的左边
	4,分别对左右子序列重复前三步
*/
func quick(arr []int, left, right int) {
	// 选取中心轴,默认选取左边第一个
	pivot := arr[left]
	// 记录中心轴的位置
	p := left
	// 定义左右指针
	i, j := left, right

	//开始遍历，如果左右指针重叠，停止遍历
	for i <= j {
		// 因此是从左边选的中心轴，所有要从右边开始遍历,右指针指向的数比中心轴大的，右指针就向左移动
		for j >= p && arr[j] >= pivot {
			j--
		}
		// 如果 不满足条件，就交换将右指针指向的数往中心轴所指向的位置赋值
		if j >= p {
			arr[p] = arr[j]
			p = j
		}
		// 判断左指针指向的数是否小于中心轴
		if i <= p && arr[i] <= pivot {
			i++
		}
		// 填空位
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
	}
	// 确定中心轴的位置
	arr[p] = pivot
	// 如果左边的元素大于1个，就调用，否则不调用
	if p-left > 1 {
		quick(arr, left, p-1)
	}
	if right-p > 1 {
		quick(arr, p+1, right)
	}

}
func QuickSort(arr []int) {
	// quickSort(arr, 0, len(arr)-1)
	quick(arr, 0, len(arr)-1)
}

func main() {

	var arr = [...]int{21, 32, 43, 6, 14, 34, 65, 90}
	fmt.Println("arr=", arr)
	bubble(&arr)
	fmt.Println("bubble arr=", arr)

	var arr1 = [...]int{21, 20, 25, 12, 89, 56, 50}
	fmt.Println("arr=", arr1)
	QuickSort(arr1[:])
	fmt.Println("quickSort arr=", arr1)

	var arr2 = [...]int{21, 20, 25, 12, 89, 56, 50}
	fmt.Println("arr=", arr2)
	QuickSort(arr2[:])
	fmt.Println("quick arr=", arr2)

	fmt.Println()
	arr3 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("arr=", arr3)
	QuickSort(arr3[:])
	fmt.Println("qq arr", arr3)

}
