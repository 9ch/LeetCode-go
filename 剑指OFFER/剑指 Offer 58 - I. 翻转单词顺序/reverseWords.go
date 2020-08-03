package main

import (
	"fmt"
	"strings"
)

/**
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

 

示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 

说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof
*/

func main() {

	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(strings.TrimSpace(" "))
}

//直接使用内建函数完成，
func reverseWords(s string) string {

	strList := strings.Split(s," ")
	var res []string
	for i :=len(strList)-1;i>=0;i--{
		str := strings.TrimSpace(strList[i])
		if  len(str)>0 {
			res = append(res,strList[i])
		}
	}
	return strings.Join(res," ")

}


//双指针.
func reverseWords2(s string) string {

	s = strings.TrimSpace(s)
	var result []byte

	j := len(s) - 1
	i := j
	temp := []byte(s)

	for i >= 1{
		for i >= 1 && temp[i] != ' ' {
			i--
		}
		result = append(result, temp[i:j+1]...)
		for temp[i] == ' ' {
			i--
		}
		j = i
	}
	return string(result)

}
