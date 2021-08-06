package main

import (
	"fmt"
	"strings"
)

/**
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


提示：你可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
*/
func main() {
	//s := "leetcode"
	//s2 := "loveleetcode"
	//fmt.Println(s)
	fmt.Println(firstUniqChar("cc"))
}

//第一次遍历，将值放入到一个map之中，并且存储他最后一次出现的位置。
//第二次遍历，如果第一次出现的位置跟最后一次出现的位置，那么就代表他只出现了一次，直接将位置返回即可，如果不相等，则直接将值改为-1（改成其他不为正整数的值也可以）
func firstUniqChar(s string) int {

	str := []byte(s)
	temp := make(map[byte]int)
	result := -1
	for k, v := range str {
		temp[v] = k
	}
	for k, v := range str {
		if temp[v] == k {
			return k
		} else {
			temp[v] = -2
		}
	}

	return result
}

// 利用唯一性求解
func firstUniqChar2(s string) int {
	for i := 0; i < len(s); i++ {
		if strings.Index(s, string(s[i])) == strings.LastIndex(s, string(s[i])) {
			return i
		}
	}
	return -1
}
