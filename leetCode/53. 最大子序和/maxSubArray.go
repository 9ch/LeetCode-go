package main

import (
	"fmt"
	"math"
)

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
*/

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(arr)
	fmt.Println(maxSubArray(arr))
}

//暴力解法，先寻找出他所有的子序列，然后再计算和， 最后再比一下大小，返回最大的和。

//动态规划，个人理解就是找规律，从最小的结果集开始寻找他们的规律。
func maxSubArray(nums []int) int {

	result := math.MinInt64
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	for _, v := range dp {
		result = max(result, v)
	}

	return result
}

func max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
