package main

import (
	"fmt"
	"strings"
)

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

 

示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."
 

限制：

0 <= s 的长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
*/
func main() {
	s := "We are happy."
	fmt.Println(replaceSpace2(s))
}

//循环遍历
func replaceSpace(s string) string {

	str := []byte(s)
	result := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			result += "%20"
		} else {
			result += string(str[i])
		}
	}
	return result
}

//内建函数
func replaceSpace2(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
