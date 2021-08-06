package main

import "fmt"

//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
//
//进阶：不要 使用任何内置的库函数，如  sqrt 。
//
//
//
//示例 1：
//
//输入：num = 16
//输出：true
//示例 2：
//
//输入：num = 14
//输出：false
//
//
//提示：
//
//1 <= num <= 2^31 - 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-perfect-square

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}

//牛顿迭代
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	x := num >> 1
	for x*x > num {
		x = (x + num/x) >> 1
	}
	return x*x == num
}
