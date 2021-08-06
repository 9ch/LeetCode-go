package main

import "fmt"

//给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
//
//例如：
//
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//
//
//示例 1：
//
//输入：columnNumber = 1
//输出："A"
//示例 2：
//
//输入：columnNumber = 28
//输出："AB"
//示例 3：
//
//输入：columnNumber = 701
//输出："ZY"
//示例 4：
//
//输入：columnNumber = 2147483647
//输出："FXSHRXW"
//
//
//提示：
//
//1 <= columnNumber <= 231 - 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/excel-sheet-column-title

func main() {
	fmt.Println(numberToTitle(1))
	fmt.Println(numberToTitle(28))
	fmt.Println(numberToTitle(701))
	fmt.Println(numberToTitle(2147483647))
}

func numberToTitle(number int) (r string) {
	result := make([]string, 0)

	for number > 0 {
		number--
		temp := number % 26
		result = append(result, string(uint8(temp)+'A'))
		number /= 26
	}
	for i := len(result) - 1; i >= 0; i-- {
		r += result[i]
	}
	return
}
