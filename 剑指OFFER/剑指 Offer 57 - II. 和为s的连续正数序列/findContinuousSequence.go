package main

import "fmt"

//输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
//
//
//
//示例 1：
//
//输入：target = 9
//输出：[[2,3,4],[4,5]]
//示例 2：
//
//输入：target = 15
//输出：[[1,2,3,4,5],[4,5,6],[7,8]]
//
//
//限制：
//
//1 <= target <= 10^5
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof

func main() {
	fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
}

//暴力循环，穷尽
func findContinuousSequence(target int) [][]int {
	var result [][]int
	i := 1
	for i < target {
		var temp []int
		count := 0
		j := i
		for count < target {
			temp = append(temp, j)
			count += j
			j++
		}
		if count == target {
			result = append(result, temp)
		}
		i++
	}
	return result
}

//滑动出窗口
