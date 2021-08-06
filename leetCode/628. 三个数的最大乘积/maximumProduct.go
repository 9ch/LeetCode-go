package main

import (
	"fmt"
	"sort"
)

//给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：6
//示例 2：
//
//输入：nums = [1,2,3,4]
//输出：24
//示例 3：
//
//输入：nums = [-1,-2,-3]
//输出：-6
//
//
//提示：
//
//3 <= nums.length <= 104
//-1000 <= nums[i] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-product-of-three-numbers

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{-1, -2, -3}
	arr4 := []int{-100, -98, -1, 2, 3, 4}
	fmt.Println(maximumProduct(arr1))
	fmt.Println(maximumProduct(arr2))
	fmt.Println(maximumProduct(arr3))
	fmt.Println(maximumProduct(arr4))
}

//排序取最大的三个值，最小的两个值即可求解
func maximumProduct(nums []int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}
	sort.Ints(nums)
	var t1, t2 int
	r1, r2 := nums[0]*nums[1], nums[length-1]*nums[length-2]
	t1 = r1 * nums[length-1]

	t2 = r2 * nums[length-3]

	if t1 > t2 {
		return t1
	} else {
		return t2
	}
}
