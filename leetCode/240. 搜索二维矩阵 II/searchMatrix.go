package main

import "fmt"

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
//
//
//示例 1：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
//输出：true
//示例 2：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
//输出：false
//
//
//提示：
//
//m == matrix.length
//n == matrix[i].length
//1 <= n, m <= 300
//-109 <= matix[i][j] <= 109
//每行的所有元素从左到右升序排列
//每列的所有元素从上到下升序排列
//-109 <= target <= 109
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii

func main() {
	//matrix := [][]int{
	//	{1,4,7,11,15},
	//	{2,5,8,12,19},
	//	{3,6,9,16,22},
	//	{10,13,14,17,24},
	//	{18,21,23,26,30},
	//}
	//fmt.Println(searchMatrix(matrix,5))
	//fmt.Println(searchMatrix(matrix,20))
	fmt.Println(searchMatrix([][]int{{-1, 3}}, 3))
}

// 搜索二维矩阵，二分查找法
func searchMatrix(matrix [][]int, target int) bool {
	left, top, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1 //先设置 4 个坐标

	for left <= right && top <= bottom {
		middle := []int{(left + right) >> 1, (top + bottom) >> 1} //求出中间节点的坐标
		if target < matrix[middle[0]][middle[1]] {
			right = middle[1] - 1
			bottom = middle[0] - 1
		} else if target > matrix[middle[0]][middle[1]] {
			left = middle[1] + 1
			top = middle[0] + 1
		} else {
			return true
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	right, bottom := len(matrix[0]), len(matrix)
	i, j := bottom-1, 0
	for i >= 0 && j < right {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return false
}
