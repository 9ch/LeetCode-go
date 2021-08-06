package main

import (
	"fmt"
	"sort"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
//示例 1:
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//示例 2:
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
//
//提示：
//
//1 <= k <= nums.length <= 104
//-104 <= nums[i] <= 104
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/xx4gT2

func main() {
	s1 := []int{3, 2, 1, 5, 6, 4}
	s2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(s1, 2))
	fmt.Println(findKthLargest(s2, 4))
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
