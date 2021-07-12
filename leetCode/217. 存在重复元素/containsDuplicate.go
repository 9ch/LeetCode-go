package main

import "fmt"

//给定一个整数数组，判断是否存在重复元素。
//
//如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
//
//
//
//示例 1:
//
//输入: [1,2,3,1]
//输出: true
//示例 2:
//
//输入: [1,2,3,4]
//输出: false
//示例 3:
//
//输入: [1,1,1,3,3,4,3,2,4,2]
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/contains-duplicate

func main() {
	arr1 := []int{1, 2, 3, 1}
	fmt.Println(containDuplicate(arr1))
	arr2 := []int{1, 2, 3, 4}
	fmt.Println(containDuplicate(arr2))
	arr3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containDuplicate(arr3))
}

func containDuplicate(nums []int) bool {
	var result bool
	mp := make(map[int]int, len(nums))

	for _, v := range nums {
		mp[v]++
	}
	for _, v := range mp {
		if v > 1 {
			return true
		}
	}

	return result
}
