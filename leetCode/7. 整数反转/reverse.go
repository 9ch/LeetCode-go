package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0


提示：

-231 <= x <= 231 - 1


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	x := -123
	fmt.Println(x)
	fmt.Println(reverse(x))
	fmt.Println(reverse1(x))
}

/**
暴力，直接变成字符串处理
*/
func reverse(x int) int {
	result := 0

	if x == 0 { //如果是 0 的话，直接返回就行
		return result
	}

	temp := strconv.Itoa(x)
	for i := len(temp) - 1; i >= 0; i-- {
		if i == 0 && string(temp[i]) == "-" { //如果首位是负数的话，则将数值直接返回
			result = 0 - result
		} else {
			s, _ := strconv.Atoi(string(temp[i]))
			result = result*10 + s
		}

	}
	if result < math.MaxInt32 && result > math.MinInt32 {
		return result
	} else {
		return 0
	}
}

/**
执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
2.2 MB
, 在所有 Go 提交中击败了
9.50%
的用户
*/

//使用对10 取模的方法，省却了变成字符串的步骤
func reverse1(x int) int {

	result := 0
	if x == 0 {
		return result
	}

	for x != 0 {
		result = result*10 + x%10
		if result > math.MaxInt32 || result < math.MinInt32 {
			result = 0
			break
		}
		x = x / 10
	}

	return result
}

/**
执行用时：
0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：
2.1 MB, 在所有 Go 提交中击败了99.24%的用户
*/
