package main

import "fmt"

/**
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
 

限制：

2 <= n <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
*/
func main() {

	//nums := []int{2, 3, 1, 0, 2, 5, 3}
	nums2 := []int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//fmt.Println(nums)
	//fmt.Println(findRepeatNumber(nums))
	fmt.Println(findRepeatNumber(nums2))
	//fmt.Println(findRepeatNumber2(nums))
	fmt.Println(findRepeatNumber2(nums2))
}

//暴力循环
func findRepeatNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i]==nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}

//哈希

func findRepeatNumber2(nums []int) int {

	 temp:= make(map[int]int,len(nums))

	for _, v := range nums {
		if temp[v] != 0 {
			return v
		}
		temp[v]++
	}
	return 0
}
