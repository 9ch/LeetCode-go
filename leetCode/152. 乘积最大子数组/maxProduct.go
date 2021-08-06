package main

import (
	"fmt"
	"math"
	"os"
)

//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//
//
//示例 1:
//
//输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//
//输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
//
//作者：力扣 (LeetCode)
//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions/xmk3rv/
//来源：力扣（LeetCode）

func main() {
	fmt.Fprintln(os.Stdout, maxProduct([]int{2, 3, -2, 4}))
	fmt.Fprintln(os.Stdout, maxProduct([]int{-2, 3, -4}))
}

// 动态规划
func maxProduct(nums []int) int {
	result := math.MinInt32
	imax, imin := 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tmp := imax
			imax = imin
			imin = tmp
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])

		result = max(result, imax)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
