package main

import "fmt"

//给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
//
//
//
//示例1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//示例 4:
//
//输入: s = ""
//输出: 0
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s由英文字母、数字、符号和空格组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

//暴力循环，从 索引 0 位置开始遍历，每次记录遇到的第一个相同元素之间的记录，放到一个数组里面，然后遍历数组取出最大值即可。
func lengthOfLongestSubstring(s string) int {

	if len(s) < 1 {
		return 0
	}

	counts := make([]int, len(s)) //定义一个切片用来存储当前不重复的最大数目

	for i := 0; i < len(s); i++ {
		m := make(map[uint8]int, len(s))
		m[s[i]] = 1
		counts[i] = 1
		for j := i + 1; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				break
			}
			counts[i]++
			m[s[j]] = 1
		}
	}
	result := counts[0]
	for _, v := range counts {
		if v > result {
			result = v
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//滑动窗口
func length2(s string) int {
	//TODO
	if len(s) < 1 {
		return 0
	}

	return 0
}
