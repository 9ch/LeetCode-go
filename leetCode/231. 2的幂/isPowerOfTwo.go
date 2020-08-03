package main

import (
	"fmt"
)

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:

输入: 1
输出: true
解释: 20 = 1
示例 2:

输入: 16
输出: true
解释: 24 = 16
示例 3:

输入: 218
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-two
*/

func main() {
	fmt.Print(2 & 1)
	fmt.Println(isPowerOfTwo(2))
}

//使用位运算通过计算1出现的次数来判断，因为2的幂次方，只有最左边的位数是1。
func isPowerOfTwo(n int) bool {

	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}

	for n > 0 {
		if n == 1 {
			return true
		}
		if n&1 == 1 {
			return false
		}
		n = n >> 1
	}
	return false
}
