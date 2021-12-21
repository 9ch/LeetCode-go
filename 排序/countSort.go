package main

import (
	"fmt"
	"math"
)

//计数排序

func main() {
	fmt.Println(countSort([]int{4, 4, 6, 5, 3, 2, 8, 1, 7, 5, 6, 0, 10}))
}

func countSort(nums []int) []int {
	max := math.MinInt32
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	temp := make([]int, max+1)

	for i := 0; i < length; i++ {
		temp[nums[i]]++
	}
	var result []int
	for i := 0; i < max+1; i++ {
		for j := 0; j < temp[i]; j++ {
			result = append(result, i)
		}
	}
	return result
}
