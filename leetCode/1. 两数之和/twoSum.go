package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	//fmt.Println(nums, target)
	fmt.Println(twoSum2(nums, target))
}

//暴力解法，直接循环比对数组元素。
func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	for k := range nums {
		for j := k + 1; j < len(nums); j++ {
			if nums[k]+nums[j] == target {
				result[0], result[1] = k, j
				return result
			}
		}
	}
	return result
}

//方法2，利用哈希存储的方式进行比对。如果已经存在了，则直接返回对应的索引位置数组
func twoSum2(nums []int, target int) []int {
	result := []int{0, 0}
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := maps[target-nums[i]]; ok {
			result[0], result[1] = v, i
			return result
		} else {
			maps[nums[i]] = i
		}
	}
	return result
}
