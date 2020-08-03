package main

import (
	"fmt"
)

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
*/

func main() {
	nums := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}

	fmt.Println(singleNumber(nums), singleNumber(nums2))
}

//位运算
func singleNumber(nums []int) int {



	//先将
	var result int = 0

	//进行位运算，拿当前值跟结果值进行异或操作。
	for _, v := range nums { //a^b^a=a^a^b=b a^a=0 a^0=a
		result = v ^ result
	}
	return result
}
