package main

import (
	"fmt"
	"math"
)

/**
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-break
*/

func main() {

	fmt.Println(integerBreak(10))
}

/**
数学思路
4: 3*1=2*2 =4
5: 3*2=6
6: 3*3=9
7: 3*3*1:3*4=12
8: 3*3*2 = 18
9: 3*3*3=27
10: 3*3*3*1 : 3*3*4 = 36
只需要统计3的个数，如果最后一个为3*1的话则将3*1补成2*2
 */
func integerBreak(n int) int {

	if n == 2 || n == 3 {
		return n - 1
	}

	num3, num2 := 0, 0

	for n > 0 {
		n -= 3
		num3++
		if n == 2 {
			num2 = 1
			break
		}
		if n == 1 {
			num2 = 2
			num3--
			break
		}
	}

	result := int(math.Pow(3.0, float64(num3)) * math.Pow(2.0, float64(num2)))
	return result

}

