package main

import (
	"fmt"
)

/**
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

 

示例 1:

输入: [3,0,1]
输出: 2
示例 2:

输入: [9,6,4,2,3,5,7,0,1]
输出: 8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number
*/

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
	fmt.Println(missingNumber([]int{0, 1}))
}

//利用求和来计算出哪个数字确实
func missingNumber(nums []int) int {

	if len(nums) == 1 {
		if nums[0] == 1 {
			return 0
		}
		if nums[0] == 0 {
			return 1
		}
	}

	total, result := 0, 0

	//使用递归
	//var sum func(n int) int
	//sum = func(n int) int {
	//	if n < 1 {
	//		return 0
	//	} else {
	//		total += n
	//		return sum(n - 1)
	//	}
	//}
	//sum(len(nums))

	//使用高斯公式
	total = (len(nums) * (len(nums) + 1)) / 2

	for _, v := range nums {
		result += v
	}
	return total - result
}

//利用位运算
func missingNumber2(nums []int) int {
	result := 0
	for i, k := range nums {
		result = result ^ k ^ i
	}
	return result ^ len(nums)
}
