package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func getMaxAndMin(arr [8]float64) (maxVal, minVal float64, maxIndex, minIndex int) {
	maxVal = arr[0]
	minVal = arr[0]
	for i := 1; i < len(arr); i++ {
		if maxVal < arr[i] {
			maxVal = arr[i]
			maxIndex = i
		}
		if minVal > arr[i] {
			minVal = arr[i]
			minIndex = i
		}
	}
	return
}
func getBestAndWorst(arr [8]float64) (bestVal, worstVal float64) {
	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	_, _, maxIndex, minIndex := getMaxAndMin(arr)
	avg := (sum - arr[maxIndex] - arr[minIndex]) / float64(len(arr))
	fmt.Printf("avg=%.2f\n", avg)
	max := math.Abs(arr[0] - avg)
	min := math.Abs(arr[0] - avg)
	fmt.Println(max)
	fmt.Println(min)
	for i := 0; ; i++ {
		if maxIndex != i && minIndex != i {
			worstVal = arr[i]
			bestVal = arr[i]
			break
		}
	}
	for i := 0; i < len(arr); i++ {
		if i == minIndex || i == maxIndex {
			continue
		}
		fmt.Printf("arr[%v]=%.2f : %.2f\n", i, arr[i], math.Abs(arr[i]-avg))
		if math.Abs(arr[i]-avg)-max > 0 {
			worstVal = arr[i]
			max = math.Abs(arr[i] - avg)
		}
		if math.Abs(arr[i]-avg)-min < 0 {
			bestVal = arr[i]
			min = math.Abs(arr[i] - avg)
		}
	}
	return
}
func main() {
	/*跳水比赛，8个评委打分，运动员的成绩是8个成绩去掉一个最高分，
	去掉一个最低分，剩下6个分数的平均分就是最后的得分。使用一维数组
	实现如下功能：
	1.把打最高和最低分的评委找出来
	2.找出最佳和最差评委，即打分与最后得分最接近的评委
	*/
	rand.Seed(time.Now().UnixNano())
	var arr [8]float64
	for i := 0; i < len(arr); i++ {
		str := fmt.Sprintf("%.2f", rand.Float64()*10)
		arr[i], _ = strconv.ParseFloat(str, 64)
	}
	fmt.Println(arr)
	maxVal, minVal, _, _ := getMaxAndMin(arr)
	bestVal, worstVal := getBestAndWorst(arr)
	fmt.Printf("maxVal=%v, minVal=%v\n", maxVal, minVal)
	fmt.Printf("bestVal=%v, worstVal=%v", bestVal, worstVal)
}
