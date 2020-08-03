package main

import (
	"fmt"
)

func main() {
	var arr = []int{5, 3, 1, 23, 56}
	// arr1 := arr[:]
	s := bubbleSort(arr)
	fmt.Println(s)
}

//冒泡排序
func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

//Output: [1 3 5 23 56]
