package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。



示例 1：

输入：x = 1, y = 4
输出：2
解释：
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
上面的箭头指出了对应二进制位不同的位置。
示例 2：

输入：x = 3, y = 1
输出：1


提示：

0 <=x, y <= 231 - 1

*/

func main() {
	x, y := 1, 4
	//fmt.Println(hammingDistance(x, y))
	fmt.Println(hammingDistance1(x, y))
}

// 普通解法，通过转换进制，并且补足0，然后按位比较
func hammingDistance(x int, y int) int {

	if x > y {
		x, y = y, x
	}
	x1 := strconv.FormatInt(int64(x), 2)
	y1 := strconv.FormatInt(int64(y), 2)

	result, i, j := 0, len(x1), len(y1)
	x1 = strings.Repeat("0", j-i) + x1 //补足前面的 0

	for k := 0; k < j; k++ {
		if x1[k] != y1[k] {
			result++
		}

	}

	return result
}

//利用异或取
func hammingDistance1(x, y int) int {
	result := 0
	temp := x ^ y
	xx := strconv.FormatInt(int64(temp), 2)
	result = strings.Count(xx, "1")
	return result
}
