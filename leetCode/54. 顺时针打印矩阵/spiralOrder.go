package main

import "fmt"

//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
//
//
//示例 1：
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//示例 2：
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//限制：
//
//0 <= matrix.length <= 100
//0 <= matrix[i].length <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof

func main() {
	s1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	s2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(s1))
	fmt.Println(spiralOrder(s2))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var result []int
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for i := l; i <= r; i++ {
			result = append(result, matrix[t][i])
		}
		t++
		if t > b {
			break
		}
		for i := t; i <= b; i++ {
			result = append(result, matrix[i][r])
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			result = append(result, matrix[b][i])
		}
		b--
		if t > b {
			break
		}
		for i := b; i >= t; i-- {
			result = append(result, matrix[i][l])
		}
		l++
		if l > r {
			break
		}

	}
	return result
}
