package main

import (
	"fmt"
	"strings"
)

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: []
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

func main() {
	arr1 := []string{"flower", "flow", "flight"}
	arr2 := []string{"dog", "racecar", "car"}
	fmt.Println(arr1, arr2)
	fmt.Println(longestCommonPrefix(arr1)) //"fl"
	fmt.Println(longestCommonPrefix(arr2)) //

}

func longestCommonPrefix(arr []string) string {

	if len(arr) < 1 {
		return ""
	}

	pre := arr[0]
	for _, v := range arr {
		if strings.Index(v, pre) != 0 {
			if len(pre) < 1 {
				return ""
			}
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}
