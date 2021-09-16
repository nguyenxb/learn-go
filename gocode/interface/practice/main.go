package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}
type Students []Student

func (stu Students) Len() int {
	return len(stu)
}
func (stu Students) Less(i, j int) bool {
	return stu[i].Score > stu[j].Score
}
func (stu Students) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

/*实现sort包中的Sort()方法接口*/
func main() {
	var stus Students
	for i := 0; i < 10; i++ {
		f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 10*rand.Float64()), 10)
		stu := Student{
			Name:  fmt.Sprintf("学生%v", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: f,
		}
		stus = append(stus, stu)
	}
	for _, v := range stus {
		fmt.Println(v)
	}
	fmt.Println()
	sort.Sort(stus)
	for _, v := range stus {
		fmt.Println(v)
	}

}
