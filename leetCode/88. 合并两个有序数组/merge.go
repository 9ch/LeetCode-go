package main

import (
	"sort"
)

//给你两个有序整数数组nums1 和 nums2，请你将 nums2 合并到nums1中，使 nums1 成为一个有序数组。
//
//初始化nums1 和 nums2 的元素数量分别为m 和 n 。你可以假设nums1 的空间大小等于m + n，这样它就有足够的空间保存来自 nums2 的元素。
//
//
//
//示例 1：
//
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//示例 2：
//
//输入：nums1 = [1], m = 1, nums2 = [], n = 0
//输出：[1]
//
//
//提示：
//
//nums1.length == m + n
//nums2.length == n
//0 <= m, n <= 200
//1 <= m + n <= 200
//-109 <= nums1[i], nums2[i] <= 109
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-sorted-array

func main() {

}

//循环，建立两个指针，通过建立一个额外数组的方式
func merge(nums1, nums2 []int, m, n int) {
	if n != 0 {
		i, j := 0, 0
		var result []int
		for i < m {
			if j < n {
				if nums2[j] <= nums1[i] {
					result = append(result, nums2[j])
					j++
				} else {
					result = append(result, nums1[i])
					i++
				}
				continue
			} else {

				break
			}
		}
		nums1 = append(result, nums2[j:]...)
	}

}

//插入，然后排序一下下
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	for m+i < m+n {
		nums1[m+i] = nums2[i]
		i++
	}
	sort.Ints(nums1)
}

//通过换位的操作
func merge3(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums2[i] > nums1[j] {

			}
		}
	}
}
