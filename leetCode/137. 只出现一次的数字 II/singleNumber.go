package main

import (
	"fmt"
)

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,3,2]
输出: 3
示例 2:

输入: [0,1,0,1,0,1,99]
输出: 99

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number-ii
*/

func main() {

	nums := []int{2, 2, 3, 2}
	nums2 := []int{0, 1, 0, 1, 0, 1, 9}

	//fmt.Println(2 ^ 2&2 ^ 3 ^ 3&3)
	fmt.Println(singleNumber(nums), singleNumber(nums2))
}

//通过map来计算
func singleNumber(nums []int) int {

	m := make(map[int]int)
	for _, k := range nums {
		m[k] += 1
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

//位运算,网上看的，没想出来这种思路。。
func singleNumber2(nums []int) int {

	var sum, result int

	for i := 0; i < 64; i++ {

		sum = 0

		for _, num := range nums {

			sum += (num >> i) & 1
		}

		result |= (sum % 3) << i
	}
	return result
}
