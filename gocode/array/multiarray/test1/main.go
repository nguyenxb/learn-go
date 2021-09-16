package main

import "fmt"

func main() {
	// 用二维数存班级的学生成绩，并统计每个班的平均分
	//和所有班级的平均分
	var scores [2][3]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入 %v 班 第 %v 个同学的成绩：\n", i, j)
			fmt.Scanln(&scores[i][j])
		}
	}

	totalSum := 0.0
	for i, v := range scores {
		sum := 0.0
		for _, v2 := range v {
			sum += v2
		}
		totalSum += sum
		fmt.Printf("%v 班的总分%v,平均分是%v\n", i, sum,
			sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分%v,平均分是%v", totalSum,
		totalSum/float64(len(scores)))
}
