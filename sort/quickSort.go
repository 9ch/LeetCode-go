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
	if length <= 1 {
		return arr
	}
	mid := arr[length/2]
	left := []int{}
	right := []int{}

	for i := 0; i < length; i++ {
		if arr[i] < mid {
			left = append(left, arr[i])
		} else {
			right = append(right, mid)
		}
	}
	left = quickSort(left)
	left = append(left, mid)
	right = quickSort(right)
	result := append(left, right...)
	return result
}
