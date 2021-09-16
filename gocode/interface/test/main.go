package main

import (
	"fmt"
	"math/rand"
	"sort" // 排序包
)

// 声明一个Hero 结构体
type Hero struct {
	Name string
	Age  int
}

// 声明一个Hero结构体的切片类型
type HeroSlice []Hero

// 对数据类型HeroSlice实现len(),Less(),Swap()方法
// 以实现Interface接口
//实现len()方法
func (Hs HeroSlice) Len() int {
	return len(Hs)
}

// Less()方法是决定你使用什么标准进行排序的
// 1.按Hero的年龄从小到大进行排序
func (Hs HeroSlice) Less(i, j int) bool {
	return Hs[i].Age < Hs[j].Age
}
func (Hs HeroSlice) Less1(i, j int) bool {
	return len(Hs[i].Name) > len(Hs[j].Name)
}

// 实现Swap方法
func (Hs HeroSlice) Swap(i, j int) {
	Hs[i], Hs[j] = Hs[j], Hs[i]
}

//

func main() {
	var intSlice = []int{0, -1, 10, 7, 9}
	// 使用系统排序包
	fmt.Println(intSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄 %v", rand.Intn(20)),
			Age:  rand.Intn(100) + 1,
		}
		// 将hero添加到切片中
		heros = append(heros, hero)
	}
	// 排序前
	for _, v := range heros {
		fmt.Println(v)
	}
	// 进行排序
	sort.Sort(heros)
	fmt.Println()
	// 排序后
	for _, v := range heros {
		fmt.Println(v)
	}

}
