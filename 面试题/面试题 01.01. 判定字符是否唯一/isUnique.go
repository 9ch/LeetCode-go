package main

import "fmt"

//实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
//
//示例 1：
//
//输入: s = "leetcode"
//输出: false
//示例 2：
//
//输入: s = "abc"
//输出: true
//限制：
//
//0 <= len(s) <= 100
//如果你不使用额外的数据结构，会很加分。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/is-unique-lcci

func main() {
	fmt.Println(isUnique2("leetcode"))
	fmt.Println(isUnique2("abc"))
}

//哈希
func isUnique(astr string) bool {
	m := make(map[int32]int, len(astr))

	for _, v := range astr {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = 1
	}

	return true
}

//双层遍历
func isUnique2(astr string) bool {
	for i := 0; i < len(astr); i++ {
		for j := i + 1; j < len(astr); j++ {
			if astr[i] == astr[j] {
				return false
			}
		}
	}
	return true
}
