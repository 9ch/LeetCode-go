package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
*/
func main() {
	str1 := "A man, a plan, a canal: Panama"
	str2 := "race a car"

	fmt.Println(str1, str2)
	fmt.Println(isPalindrome(str1))
}

func isPalindrome(str string) bool {

	//先把字符变成小写。
	str = strings.ToLower(str)
	i, j := 0, len(str)-1
	for i < j {
		//如果字符不是数字或者字母，则将相应的索引移位，并跳出此次循环，继续匹配。
		if !(unicode.IsNumber(rune(str[i])) || unicode.IsLetter(rune(str[i]))) {
			i++
			continue
		}
		if !(unicode.IsNumber(rune(str[j])) || unicode.IsLetter(rune(str[j]))) {
			j--
			continue
		}
		//如果两个位置的字符不一致，则直接跳出循环
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome2(s string) bool {
	str := strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !((unicode.IsNumber(rune(str[i]))) || unicode.IsLetter(rune(str[i]))) {
			i++
			continue
		}
		if !((unicode.IsNumber(rune(str[j]))) || unicode.IsLetter(rune(str[j]))) {
			j--
			continue
		}
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
