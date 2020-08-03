package main

import "fmt"

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
*/
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(nums)
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	var result [][]int
	if length < 3 {
		return result
	}

	n := 0
	for i := 0; i < length; i++ {
		for j := 1; j < length; j++ {
			for k := 2; k < length; k++ {
				if (nums[i] + nums[j] + nums[k]) == 0 {
					result[n][0] = i
					result[n][1] = j
					result[n][2] = k
					n++
				}
			}
		}
	}

	return result
}
