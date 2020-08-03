package main

import (
	"fmt"
	"strings"
)

/**

实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr/
*/

func main() {
	haystack1 := "hello"
	needle1 := "ll"

	haystack2 := "aaaaa"
	needle2 := "bba"

	fmt.Println(haystack1, needle1, haystack2, needle2)
	fmt.Println(strStr(haystack1, needle1))
}

//解题思路：Sunday 匹配（参考）
//Sunday 算法是 Daniel M.Sunday 于1990年提出的字符串模式匹配。其核心思想是：在匹配过程中，模式串发现不匹配时，算法能跳过尽可能多的字符以进行下一步的匹配，从而提高了匹配效率。
func strStr(haystack, needle string) int {

	haystackLength := len(haystack)
	needleLength := len(needle)
	if haystackLength < needleLength {
		return -1
	}
	hayIndex, needleIndex := 0, 0

	//设置循环次数
	for needleIndex < needleLength {

		//当主字符串所以超过字符串长度时，则直接返回，因为没有查询到。
		if hayIndex > haystackLength-1 {
			return -1
		}

		//当目标和当前字符串的索引位置的值相等时，则将索引值+1
		if haystack[hayIndex] == needle[needleIndex] {
			hayIndex++
			needleIndex++
		} else {
			//设置每次需要挪动的距离。
			next := hayIndex - needleIndex + needleLength //(核心)
			//移动的距离小于目标字符串的时候，说明没有查找到字符，则直接返回-1
			if next < haystackLength {
				//计算needle字符串中最后出现的索引位置，如果存在，则将needle向前挪动next-step位，如果不存在，则将步进+1
				step := strings.LastIndex(needle, string(haystack[next])) //(核心)
				if step == -1 {
					hayIndex = next + 1 //(核心)
				} else {
					hayIndex = next - step //(核心)
				}
				needleIndex = 0//每次needle的索引值都从0开始。
			} else {
				return -1
			}
		}
	}

	return hayIndex - needleIndex

}
