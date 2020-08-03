package main

import "fmt"

/**
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
*/

func main() {

	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums2 := []int{4, 10, 4, 3, 8, 9}
	//fmt.Println(nums[:3])
	fmt.Println(lengthOfLIS(nums), lengthOfLIS(nums2),lengthOfLIS([]int{1, 9, 5, 9, 3}))
	fmt.Println(lengthOfLIS([]int{0,0}))
}

//动态规划。又名：找规律
func lengthOfLIS(nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	dp := make([]int, len(nums)) //定义一个存储长度的切片
	result :=1

	for i := 0; i < len(nums); i++ {
		dp[i]=1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
