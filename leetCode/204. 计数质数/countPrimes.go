package main

import (
	"fmt"
)

//统计所有小于非负整数 n 的质数的数量。
//
//
//
//示例 1：
//
//输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//示例 2：
//
//输入：n = 0
//输出：0
//示例 3：
//
//输入：n = 1
//输出：0
//
//
//提示：
//
//0 <= n <= 5 * 106
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-primes

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	count := 1
	for i := 3; i < n; i += 2 {
		if isPrimes(i) {
			count++
		}
	}
	return count
}

func isPrimes(i int) bool {
	for j := 3; j < i; j += 2 {
		if i%j == 0 {
			return false
		}
	}
	return true
}
