package main

import (
	"fmt"
	"strings"
)

//给定一个单词，你需要判断单词的大写使用是否正确。
//
//我们定义，在以下情况时，单词的大写用法是正确的：
//
//全部字母都是大写，比如"USA"。
//单词中所有字母都不是大写，比如"leetcode"。
//如果单词不只含有一个字母，只有首字母大写， 比如 "Google"。
//否则，我们定义这个单词没有正确使用大写字母。
//
//示例 1:
//
//输入: "USA"
//输出: True
//示例 2:
//
//输入: "FlaG"
//输出: False
//注意: 输入是由大写和小写拉丁字母组成的非空单词。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/detect-capital

func main() {
	fmt.Println(detectCapitaUse("USA"))
	fmt.Println(detectCapitaUse("Flag"))
	fmt.Println(detectCapitaUse(""))
}

func detectCapitaUse(word string) bool {
	if word == strings.ToLower(word) || word == strings.ToUpper(word) {
		return true
	}
	for i := 1; i < len(word); i++ {
		if string(word[i]) != strings.ToLower(string(word[i])) {
			return false
		}
	}
	return true
}
