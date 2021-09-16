package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	address string
}

// 多重继承// 尽量不使用多重继承
type TV struct {
	Goods
	Brand
	// Brand
	int
	// int // 最多只能有一个同类型的匿名字段
}
type TV2 struct {
	*Goods
	*Brand
}

func main() {
	/*
		嵌套结构体后，也可以直接给各个结构体的字段赋值
		直接指定各个匿名结构体字段的值
	*/
	tv := TV{Goods{"123", 12.5}, Brand{"adsa", "北京"}, 3}
	fmt.Println("tv", tv)
	tv2 := TV{
		Goods{
			Name:  "1231",
			Price: 12.51,
		},
		Brand{
			Name:    "adsa1",
			address: "北京1",
		},
		// 初始化匿名字段
		1,
	}
	fmt.Println("tv2=", tv2)
	tv3 := TV2{
		&Goods{
			Name:  "1231",
			Price: 12.51,
		},
		&Brand{
			Name:    "adsa1",
			address: "北京1",
		},
	}
	fmt.Println("tv3=", tv3)
	fmt.Println("tv3 goos=", *tv3.Goods, "tv3 brand", tv3.Brand)

}
