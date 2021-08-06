package main

import (
	"fmt"
	"sort"
)

//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
//示例 1：
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//示例 2：
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//提示：
//
//0 <= nums.length <= 105
//-109 <= nums[i] <= 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-consecutive-sequence

func main() {
	s1 := []int{100, 4, 200, 1, 3, 2}
	s2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(s1))
	fmt.Println(longestConsecutive(s2))
}

//动态规划，超过 ON 复杂度
func longestConsecutive2(nums []int) int {
	sort.Ints(nums)
	//TODO
	return 0
}

// 哈希表实现。
func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	longestStreak := 0

	for num := range m {
		if !m[num-1] {
			currentNum := num + 1
			currentStreak := 1
			for m[currentNum] {
				currentStreak++
				currentNum++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
