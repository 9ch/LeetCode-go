package main

import "fmt"

//给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
//
//示例 1 :
//
//输入: 2736
//输出: 7236
//解释: 交换数字2和数字7。
//示例 2 :
//
//输入: 9973
//输出: 9973
//解释: 不需要交换。
//注意:
//
//给定数字的范围是 [0, 108]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-swap

func main() {
	fmt.Println(maximumswap(2736))
	fmt.Println(maximumswap(9973))
	fmt.Println(maximumswap(98368))
	fmt.Println(maximumswap(10))

}

// 穷举出来所有的数字，并记录最大值，返回即可
func maximumswap(num int) int {

	if num < 10 {
		return num
	}
	var t []int
	max := num //记录最大值
	for num > 0 {
		t = append(t, num%10)
		num /= 10
	}
	for i := len(t) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			t[i], t[j] = t[j], t[i]
			if s := SliceToNum(t); s > max {
				max = s
			}
			t[i], t[j] = t[j], t[i]
		}
	}
	return max
}

func SliceToNum(num []int) (result int) {
	for i := len(num) - 1; i >= 0; i-- {
		result = result*10 + num[i]
	}
	return
}
