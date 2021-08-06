package main

import (
	"fmt"
)

//	选择排序
// 时间复杂度 O(n^2),空间复杂度 O(1)
// 第一次,先找出最小(大)的值,放到第一的位置
// 第二次,找出最小(大)的值,放到第二的位置

func main() {
	arr := []int{-1, 7, 8, 2, 1, 3, 4, 6, 5}
	fmt.Println(arr)
	fmt.Println(selection(arr))
}

func selection(nums []int) []int {
	length := len(nums)

	if length < 2 {
		return nums
	}
	var min, k int
	for i := 0; i < length; i++ {

		for j := i; j < length; j++ {
			if j == i || nums[j] < min {
				min = nums[j]
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	return nums
}
