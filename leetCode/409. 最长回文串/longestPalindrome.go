package main

import "fmt"

//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
//在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
//注意:
//假设字符串的长度不会超过 1010。
//
//示例 1:
//
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindrome

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}

func longestPalindrome(s string) int {
	m := make(map[uint8]int)
	result := 0
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]] == 2 {
			result += 2
			m[s[i]] = 0
		}
	}
	fmt.Println(m)
	if result < len(s) {
		result += 1
	}
	return result
}
