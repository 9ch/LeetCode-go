package main

import "fmt"

//和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
//
//现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。
//
//数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
//
//
//
//示例 1：
//
//输入：nums = [1,3,2,2,5,2,3,7]
//输出：5
//解释：最长的和谐子序列是 [3,2,2,2,3]
//示例 2：
//
//输入：nums = [1,2,3,4]
//输出：2
//示例 3：
//
//输入：nums = [1,1,1,1]
//输出：0
//
//
//提示：
//
//1 <= nums.length <= 2 * 104
//-109 <= nums[i] <= 109
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-harmonious-subsequence

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(findLHS([]int{1, 2, 3, 4}))
	fmt.Println(findLHS([]int{1, 1, 1, 1}))
}

// 利用 map，计算相邻元素的最大值即可
func findLHS(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}
	set := make(map[int]int)
	for _, v := range nums {
		set[v]++
	}
	result := 0 //记录最大值
	for k, v := range set {
		var left, right int
		if vv, ok := set[k-1]; ok {
			left = vv + v //当左边有值的时候，则直接将两边相加就可。
		}
		if vv, ok := set[k+1]; ok {
			right = vv + v // 当右边有值的时候
		}
		result = max(result, max(left, right)) // 比较左右两边的最大值
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
