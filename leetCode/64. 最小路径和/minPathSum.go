package main

import "fmt"

/**
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
https://leetcode-cn.com/problems/minimum-path-sum/
*/

func main() {

	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	fmt.Println(grid)
	fmt.Println(minPathSum(grid))

}


//动态规划，需要加深印象，主要是为了寻找他们之间的关系及状态转移方程。
func minPathSum(grid [][]int) int {

	if len(grid) <1 {
		return -1
	}


	for i := 0; i < len(grid); i++ { //循环M层，即高度，Y轴
		for j := 0; j < len(grid[0]); j++ { //循环N层，及宽度，X轴
			if i == 0 && j == 0 { //如果再原点位置，则直接跳过，不需要计算。
				continue
			}
			if i == 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1] //在上边界时
			} else if j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j] //在左边界时
			} else {
				//当前的路径最小值就是左边或者上面的最小值加当前值
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}

		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
