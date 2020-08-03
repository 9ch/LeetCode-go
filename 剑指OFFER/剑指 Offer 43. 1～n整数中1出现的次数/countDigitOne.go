package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

 

示例 1：

输入：n = 12
输出：5
示例 2：

输入：n = 13
输出：6
 

限制：

1 <= n < 2^31

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof
*/

func main() {
	fmt.Println(countDigitOne(199))
}

//超出时间限制
func countDigitOne(n int) int {
	if n == 1 {
		return n
	}
	count := 0
	for i := 1; i <= n; i++ {
		count += strings.Count(strconv.Itoa(i), "1")
	}
	return count
}
