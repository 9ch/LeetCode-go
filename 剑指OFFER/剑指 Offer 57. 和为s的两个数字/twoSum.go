package main

import "fmt"

//输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
//
//
//
//示例 1：
//
//输入：nums = [2,7,11,15], target = 9
//输出：[2,7] 或者 [7,2]
//示例 2：
//
//输入：nums = [10,26,30,31,47,60], target = 40
//输出：[10,30] 或者 [30,10]
//
//
//限制：
//
//1 <= nums.length <= 10^5
//1 <= nums[i] <= 10^6

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(arr)
	fmt.Println(twoSum(arr, 9))
	fmt.Println(twoSum2(arr, 9))
}

//哈希法
func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		if _, ok := m[target-v]; ok {
			result[0] = target - v
			result[1] = v
			return result
		}
		m[v] = k
	}

	return result
}

//暴力法˚
func twoSum2(nums []int, target int) []int {
	result := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = nums[i]
				result[1] = nums[j]
				break
			}
		}
	}

	return result
}
