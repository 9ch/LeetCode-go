package main

import (
	"fmt"
	"strings"
)

//字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
//
//示例1:
//
// 输入：s1 = "waterbottle", s2 = "erbottlewat"
// 输出：True
//示例2:
//
// 输入：s1 = "aa", s2 = "aba"
// 输出：False
//提示：
//
//字符串长度在[0, 100000]范围内。
//说明:
//
//你能只调用一次检查子串的方法吗？
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-rotation-lcci

func main() {
	fmt.Println(isFlipedString("waterbottle", "erbottlewat"))
	fmt.Println(isFlipedString("aa", "aba"))
}

//将s1 相加，如果 s2 在是反转的话，则必然存在s1里面
func isFlipedString(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	temp := s1 + s1
	i := strings.Index(temp, s2)
	if i != -1 {
		return true
	} else {
		return false
	}
}
