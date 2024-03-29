package main

import (
	"fmt"
	"strconv"
)

//字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
//
//示例1:
//
// 输入："aabcccccaaa"
// 输出："a2b1c5a3"
//示例2:
//
// 输入："abbccd"
// 输出："abbccd"
// 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
//提示：
//
//字符串长度在[0, 50000]范围内。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/compress-string-lcci

func main() {
	s1 := "aabcccccaaa"
	s2 := "abbccd"
	fmt.Println(compressString(s1))
	fmt.Println(compressString(s2))
}

//使用数组计算
func compressString(S string) string {
	if len(S) < 2 {
		return S
	}
	result := ""
	count := 1 //记录每次遍历的
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			count++
		} else {
			result = result + string(S[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}
	result = result + string(S[len(S)-1]) + strconv.Itoa(count) //防止最后一个没记录进去
	if len(result) > len(S) {
		return S
	}
	return result
}
