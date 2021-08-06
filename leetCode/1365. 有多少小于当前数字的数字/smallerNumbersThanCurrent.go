package main

import (
	"fmt"
	"sort"
)

//给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
//
//换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
//
//以数组形式返回答案。
//
//
//
//示例 1：
//
//输入：nums = [8,1,2,2,3]
//输出：[4,0,1,1,3]
//解释：
//对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。
//对于 nums[1]=1 不存在比它小的数字。
//对于 nums[2]=2 存在一个比它小的数字：（1）。
//对于 nums[3]=2 存在一个比它小的数字：（1）。
//对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。
//示例 2：
//
//输入：nums = [6,5,4,8]
//输出：[2,1,0,3]
//示例 3：
//
//输入：nums = [7,7,7,7]
//输出：[0,0,0,0]
//
//
//提示：
//
//2 <= nums.length <= 500
//0 <= nums[i] <= 100
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number

func main() {
	fmt.Println(smallerNumbersThanCurrent2([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent2([]int{6, 5, 4, 8}))
}

//暴力
func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				temp++
			}
		}
		result[i] = temp
	}
	return result
}

//排序+哈希
func smallerNumbersThanCurrent2(nums []int) []int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	var result []int
	m := make(map[int]int, len(nums))
	for i := 0; i < len(temp); i++ {
		if _, ok := m[temp[i]]; !ok {
			m[temp[i]] = i
		}
	}

	for i := 0; i < len(nums); i++ {
		result = append(result, m[nums[i]])
	}
	return result
}
