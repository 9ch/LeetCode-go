package main

import "fmt"

/**
魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。
给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。
若有多个魔术索引，返回索引值最小的一个。

示例1:

 输入：nums = [0, 2, 3, 4, 5]
 输出：0
 说明: 0下标的元素为0
示例2:

 输入：nums = [1, 1, 1]
 输出：1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/magic-index-lcci
*/

func main() {
	nums := []int{-294354269, -168926144, -62738268, 3}
	fmt.Println(findMagicIndex(nums))
}

//暴力循环
func findMagicIndex(nums []int) int {
	if len(nums) < 1 {
		return -1
	}

	for i := 0; i < len(nums)-1; i++ {
		if i == nums[i] {
			return i
		}
	}
	return -1
}

//二分查找
func findMagicIndex2(nums []int) int {

}
