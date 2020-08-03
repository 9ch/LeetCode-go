package main

import (
	"fmt"
	"strings"
)

/**
给定两个字符串, A 和 B。

A 的旋转操作就是将 A 最左边的字符移动到最右边。 例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。如果在若干次旋转操作之后，A 能变成B，那么返回True。

示例 1:
输入: A = 'abcde', B = 'cdeab'
输出: true

示例 2:
输入: A = 'abcde', B = 'abced'
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-string
*/
func main() {
	a := "abcde"
	b := "cdeab"
	c := "abced"
	//fmt.Println(a, b)
	//fmt.Println(rotateString(a, b))
	//fmt.Println(rotateString(a, c))
	fmt.Println(rotateString1(a, c))
	fmt.Println(rotateString1(a, b))
}

/**
暴力遍历，复杂度O(n)，
先将源字符串转换成字节，然后通过比对源字符串与目标字符串是否一致，如果不一致则继续重组。。
 */
func rotateString(a string, b string) bool {

	if len(a) != len(b){
		return false
	}
	if a == b {
		return true
	}
	bA := []byte(a)
	for i := 0; i < len(a); i++ {
		if string(bA) != b {
			bA = append(bA[1:], bA[:1]...) //在原数组上面重组，来节省空间，虽然这道题最大字符长度只有100...
		} else {
			return true
		}
	}
	return false
}


/**
b总是存在于a+a的子串之中。。
 */
func rotateString1(a string, b string) bool {
	totalStr := a+a
	fmt.Println(totalStr)
	return strings.Contains(totalStr,b)
}