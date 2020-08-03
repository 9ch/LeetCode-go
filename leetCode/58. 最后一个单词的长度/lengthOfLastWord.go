package main

import (
	"fmt"
	"strings"
)

/**

给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。



示例:

输入: "Hello World"
输出: 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/length-of-last-word/
*/
func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {

	s = strings.TrimRight(s, " ") //去除尾部可能有的空格
	i := len(s) - 1
	count := 0

	//从尾部开始计算，遇到第一个空格，则退出循环，否则的话，移动指针，并且累加
	for i >= 0 {
		if s[i] != ' ' {
			count++
			i--
		}else {
			break
		}
	}
	return count
}
