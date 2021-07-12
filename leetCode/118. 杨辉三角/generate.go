package main

import "fmt"

//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
//
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 5
//输出:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/pascals-triangle

func main() {
	fmt.Println(generate(5))
}

func generate(n int) [][]int {

	if n == 0 {
		return [][]int{}
	}

	result := make([][]int, n)
	result[0] = []int{1}
	if n == 1 {
		return result
	}
	result[1] = []int{1, 1}
	if n == 2 {
		return result
	}
	i := 2
	for i < n {
		temp := []int{1}
		for j := 0; j+1 < len(result[i-1]); j++ {
			temp = append(temp, result[i-1][j]+result[i-1][j+1])
		}
		temp = append(temp, 1)
		result[i] = temp
		i++
	}
	return result
}
