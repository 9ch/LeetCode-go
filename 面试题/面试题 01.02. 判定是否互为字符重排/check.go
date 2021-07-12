package main

import "fmt"

//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
//
//示例 1：
//
//输入: s1 = "abc", s2 = "bca"
//输出: true
//示例 2：
//
//输入: s1 = "abc", s2 = "bad"
//输出: false
//说明：
//
//0 <= len(s1) <= 100
//0 <= len(s2) <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/check-permutation-lcci

func main() {
	fmt.Println(check("abc", "bca"))
	fmt.Println(check("abc", "bad"))
}

//哈希的方式
func check(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[int32]int, len(s1))
	m2 := make(map[int32]int, len(s2))
	for _, v := range s1 {
		m1[v]++
	}
	for _, v := range s2 {
		m2[v]++
	}
	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}
	return true
}
