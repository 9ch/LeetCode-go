package main

import (
	"fmt"
)

//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
//输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//提示：
//
//每个字符串仅由字符 '0' 或 '1' 组成。
//1 <= a.length, b.length <= 10^4
//字符串如果不是 "0" ，就都不含前导零。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-binary

func main() {
	fmt.Println("11", "1")
	fmt.Println("1010", "1011")
}

//模拟进位
func addBinary(a, b string) string {

}
