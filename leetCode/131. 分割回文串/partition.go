package main

import (
	"fmt"
	"os"
)

//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
//回文串 是正着读和反着读都一样的字符串。
//
// 
//
//示例 1：
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//示例 2：
//
//输入：s = "a"
//输出：[["a"]]
// 
//
//提示：
//
//1 <= s.length <= 16
//s 仅由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindrome-partitioning

func main() {
	fmt.Fprint(os.Stdin, reverse("hello"))
}

func partition(s string) [][]string {
	result := make([][]string)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if
		}
	}
}

func reverse(s string) string {
	t := []byte(s)
	for i := 0; i < len(t)>>1; i++ {
		t[i], t[len(t)-1-i] = t[len(t)-1-i], t[i]
	}
	return string(t)
}
