package main

import (
	"fmt"
	"sort"
)

/**
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

 

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
 

提示：

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
*/

func main() {
	nums := []int{2,1}
	fmt.Println(firstMissingPositive(nums))
}

//暴力
func firstMissingPositive(nums []int) int {

	sort.Ints(nums)
	unique(nums)
	fmt.Println(nums)
	result := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		if nums[i] != result {
			return result
		}
		result++
	}
	fmt.Println(result)
	return result
}

//去重
func unique(nums []int) {
	sort.Ints(nums)
	for i := 1; i < len(nums); {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
}
