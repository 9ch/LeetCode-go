package main

import "fmt"

//不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
//
//示例 1:
//
//输入: a = 1, b = 2
//输出: 3
//示例 2:
//
//输入: a = -2, b = 3
//输出: 1
//通过次数52,717提交次数91,029
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-two-integers

func main() {
	fmt.Println(1, 2)
	fmt.Println(-2, 3)
}

func getSum(a int, b int) int {
	// 异或+与运算:时间复杂度O(logSum) | 空间复杂度O(1)
	a, b = a^b, (a&b)<<1
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a

}
