package main

import "fmt"

//整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。
//
//示例1:
//
// 输入：A = 29 （或者0b11101）, B = 15（或者0b01111）
// 输出：2
//示例2:
//
// 输入：A = 1，B = 2
// 输出：2
//提示:
//
//A，B范围在[-2147483648, 2147483647]之间
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/convert-integer-lcci

func main() {
	fmt.Println(convertInteger(29, 15))
}

//才进行异或比较，最后再计算 1 的个数
func convertInteger(A int, B int) int {
	result := A ^ B
	count := 0
	for i := 0; i < 32; i++ {
		count += result >> i & 1
	}
	return count
}
