package main

import "fmt"

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//你可以假设数组中无重复元素。
//
//示例 1:
//
//输入: [1,3,5,6], 5
//输出: 2
//示例 2:
//
//输入: [1,3,5,6], 2
//输出: 1
//示例 3:
//
//输入: [1,3,5,6], 7
//输出: 4
//示例 4:
//
//输入: [1,3,5,6], 0
//输出: 0
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-insert-position

func main() {
	arr1 := []int{1, 3, 5, 6}
	//fmt.Println(searchInsert(arr1, 5))
	//fmt.Println(searchInsert(arr1, 2))
	fmt.Println(searchInsert(arr1, 7))
	fmt.Println(searchInsert(arr1, 0))
}

//二分查找法
func searchInsert(nums []int, target int) int {

	length := len(nums)
	//如果长度小于 1，或者目标为 0，必然会在数组的第一个位置
	if length == 0 || target == 0 {
		return 0
	}

	mid := length / 2   //去一个中间值
	l, r := 0, length-1 //定义左右边界值

	for l <= r {

		if nums[mid] == target { //如果再当前位置则直接返回
			return mid
		} else if nums[mid] < target {
			l = mid + 1
			mid = (l + r) / 2
		} else {
			r = mid - 1
			mid = (l + r) / 2
		}

	}

	return l //当左边界和右边界相等的时候，则该值在他们的下一个节点
}
