package main

import "fmt"

//给定两个数组，编写一个函数来计算它们的交集。
//
//
//
//示例 1：
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//示例 2：
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//
//
//说明：
//
//输出结果中的每个元素一定是唯一的。
//我们可以不考虑输出结果的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-arrays

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersection(nums1, nums2 []int) []int {
	if len(nums2) == 0 || len(nums2) == 0 {
		return []int{}
	}
	m := make(map[int]int)
	m1 := make(map[int]int)
	var result []int
	if len(nums1) < len(nums1) {
		for i := 0; i < len(nums1); i++ {
			m[nums1[i]] = i
		}
		for i := 0; i < len(nums2); i++ {
			if _, ok := m[nums2[i]]; ok {
				m1[nums2[i]] = i
			}
		}
	} else {
		for i := 0; i < len(nums2); i++ {
			m[nums2[i]] = i
		}
		for i := 0; i < len(nums1); i++ {
			if _, ok := m[nums1[i]]; ok {
				m1[nums1[i]] = i
			}
		}
	}
	for k, _ := range m1 {
		result = append(result, k)
	}
	return result
}
