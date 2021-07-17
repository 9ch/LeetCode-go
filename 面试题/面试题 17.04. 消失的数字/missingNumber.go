package main

import (
	"fmt"
	"sort"
)

//数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？
//
//注意：本题相对书上原题稍作改动
//
//示例 1：
//
//输入：[3,0,1]
//输出：2
//
//
//示例 2：
//
//输入：[9,6,4,2,3,5,7,0,1]
//输出：8
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/missing-number-lcci

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

//先排序，再根据索引值比对
func missingNumber2(nums []int) int {

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}

//异或
func missingNumber(nums []int) int {
	num := 0
	for i := 1; i <= len(nums); i++ {
		num ^= i
		num ^= nums[i-1]
	}
	return num
}
