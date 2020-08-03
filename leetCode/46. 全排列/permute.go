package main

import "fmt"

/**
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
*/
func main() {
	a := []int{5,4,6,2}

	//fmt.Println(a)
	fmt.Println(permute(a))
}

var result [][]int

func permute(nums []int) [][]int {

	dfs(nums, []int{})
	return result

}

func dfs(nums []int, temp []int) {

	if len(temp) == len(nums) {
		result = append(result, temp)

	} else {
		for _, v := range nums {
			if !inArray(temp, v) {
				temp = append(temp, v)
				dfs(nums, temp)
				temp = temp[:len(temp)-1]
			}
		}
	}
}

func inArray(arr []int, num int) bool {

	maps := make(map[int]interface{})

	for _, v := range arr {
		maps[v] = struct {}{}
	}
	if _, ok := maps[num]; ok {
		return true
	} else {
		return false
	}
}
