package main

import (
	"fmt"
	"sort"
)

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
	nums := []int{0, 0, 0, 0}

	fmt.Println(nums)
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)                  //先进行排序
	for i := 0; i < len(nums); i++ { //从第一个元素开始
		j, k := i+1, len(nums)-1
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] { //跳过重复值，因为已经计算过了。
			continue
		}
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				num := []int{nums[i], nums[j], nums[k]}
				res = append(res, num)
				for j < k && nums[j] == nums[j+1] { //执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,RL,R 移到下一位置，寻找新的解
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
