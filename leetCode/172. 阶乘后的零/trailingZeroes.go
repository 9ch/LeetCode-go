package main

import (
	"fmt"
)

//给定一个整数 n，返回 n! 结果尾数中零的数量。
//
//示例 1:
//
//输入: 3
//输出: 0
//解释:3! = 6, 尾数中没有零。
//示例2:
//
//输入: 5
//输出: 1
//解释:5! = 120, 尾数中有 1 个零.
//说明: 你算法的时间复杂度应为O(logn)。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/factorial-trailing-zeroes

func main() {
	fmt.Println(tailingZeroes2(30))
	fmt.Println(tailingTour(30))
}

//找出负载因子，只要有5，那么一定会出现尾数为 0 的情况
func tailingZeroes2(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n = n / 5
	}
	return count
}

func tailingTour(n int) int {
	count := 0
	for n > 0 {
		if n%5 == 0 {
			count++
		}
		n--
	}
	return count
}
