package main

import "fmt"

func main() {
	var arr = []int{5, 3, 1, 23, 56}
	s := quickSort(arr)
	fmt.Println(s)
}

//快速排序
func quickSort(arr []int) []int {

	length := len(arr)

	if length <= 1 { //递归返回条件
		return arr
	}
	mid := arr[0]         //取一个中间值
	var left, right []int //定义两个分别存储最小值和最大值的切片
	for i := 1; i < length; i++ {
		if arr[i] < mid {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = quickSort(left)
	left = append(left, mid)
	right = quickSort(right)
	return append(left, right...)
}
