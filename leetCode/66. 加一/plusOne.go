package main

import "fmt"

/**
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
*/
func main() {
	digits1 := []int{9, 9, 9}
	digits2 := []int{4, 3, 2, 1}

	fmt.Println(digits1, digits2)
	fmt.Println(plusOne(digits1), plusOne(digits2))
}

/**
题解：
需要考虑两点：
1. 正常数组[1,3,4]这种，没有进位操作，+1之后，直接返回即可
2. 需要进位的数组，类似于[9,9,9]这种，每个位置加1之后，变成了1000，则需要进行进位。
先将当前位的值进行+1操作，如果当前值对10取余之后，不为0，则代表不需要进位操作，直接返回即可。
如果没有返回，则代表这个数字有进位操作
*/
func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		//先将当前位的值进行+1操作，如果当前值对10取余之后，不为0，则代表不需要进位操作，直接返回即可。
		digits[i]++
		digits[i] = digits[i] % 10

		if digits[i] != 0 {
			return digits
		}
	}
	//如果没有返回，则代表这个数字有进位操作，类似于999之类的数字，变成四位。则需要在首位进行填位1操作。

	return append([]int{1}, digits...)
}
