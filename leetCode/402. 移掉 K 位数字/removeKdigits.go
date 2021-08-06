package main

import (
	"fmt"
)

//给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。
//
//
//示例 1 ：
//
//输入：num = "1432219", k = 3
//输出："1219"
//解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。
//示例 2 ：
//
//输入：num = "10200", k = 1
//输出："200"
//解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
//示例 3 ：
//
//输入：num = "10", k = 2
//输出："0"
//解释：从原数字移除所有的数字，剩余为空就是 0 。
//
//
//提示：
//
//1 <= k <= num.length <= 105
//num 仅由若干位数字（0 - 9）组成
//除了 0 本身之外，num 不含任何前导零
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-k-digits

func main() {
	fmt.Println(removeKdigits2("1432219", 3))
	fmt.Println(removeKdigits2("10200", 1))
	fmt.Println(removeKdigits2("10", 10))
	fmt.Println(removeKdigits2("10001", 4))
}

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}

	stack := make([]rune, 0) // 定义一个栈

	for _, v := range num {

		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > v { //比较当前元素与栈顶的元素，如果比栈顶的元素小，则不断的删除栈顶的元素
			stack = stack[0 : len(stack)-1]
			k--
		}

		if v != '0' || len(stack) != 0 { //防止第一个元素为 0
			stack = append(stack, v)
		}
	}

	for k > 0 { // k 还大于 0，则需要不断的移除栈顶的元素
		stack = stack[0 : len(stack)-1]
		k--
	}
	if len(stack) == 0 { // 防止返回的是字符串的空值“”
		return "0"
	}

	return string(stack)

}

func removeKdigits2(num string, k int) string {
	if len(num) <= k {
		return "0"
	}

	stack := make([]rune, 0)

	for _, v := range num {
		//if i == 0 && num[i] == 0 {
		//	continue
		//}
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > v {
			stack = stack[:len(stack)-1]
			k--
		}

		if v != '0' || len(stack) != 0 {
			stack = append(stack, v)
		}
	}

	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
