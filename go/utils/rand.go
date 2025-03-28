package utils

import "math/rand"

// 生成指定数量顺序的列表
func GenOrderIntList(num int) []int {
	arr := []int{}
	for i := 0; i < num; i++ {
		arr = append(arr, i+1)
	}
	return arr
}

// 生成指定数量乱序的列表
func GenDisOrderIntList(num int) []int {
	arr := []int{}
	for i := 0; i < num; i++ {
		arr = append(arr, rand.Intn(10))
	}
	return arr
}
