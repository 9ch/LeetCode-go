package main

import (
	"fmt"
)

//在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//
//示例:
//
//s = "abaccdeff"
//返回 "b"
//
//s = ""
//返回 " "
//
//
//限制：
//
//0 <= s 的长度 <= 50000

func main() {
	s := "abaccdeff"
	fmt.Println(s)
	fmt.Println(firstUniqChar(s))
}

//存在 hash 里面，再根据字符的顺序去遍历哈希的值就知道第一个了。
func firstUniqChar(s string) byte {

	m := make(map[byte]int, len(s))
	sb := []byte(s)
	for _, v := range sb {
		m[v]++
	}

	for _, v := range sb {
		if m[v] == 1 {
			return v
		}
	}

	return ' '
}

//先将字符串排序，如果当前字符与左右两边都不相等，则说明这个字符就是第一个唯一字符
//func fisrtUniqChar2(s string) byte {
//
//}
