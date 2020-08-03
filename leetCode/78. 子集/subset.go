package main

import "fmt"

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [1],
  [2],
  [3],
  [1,2],
  [1,3],
  [2,3],
  [1,2,3],
  []
]
通过次数113,114提交次数145,660

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
*/

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(subsets(nums))
}

/**
leetcode运行不通过，但是结果是对的。
主要是通过动态规划的思想解决的
[1] = [],[1]
[1,2] = 1,2,[1,2]
[1,2,3] = 1,2,3,[1,3],[2,3],[1,2,3],[1,2]
可以发现，i是i-1的所有集合元素加上i，并且多出来了0到i的所有元素和上一个最后的集合元素。

[1,2]的时候：[[1 2] [2] [1] []]
[1,2,3]的时候： [[1 2 3] [2 3] [1 3] [3] [1 3] [1] [2] [] [1 2]]
[1,2,3,4][[1 2 3 4] [2 3 4] [1 3 4] [3 4] [1 3 4] [1 4] [2 4] [4] [1 2 4] [1] [2] [3] [] [1 2 3]]

*/
func subsets(nums []int) [][]int {

	var result [][]int
	result = append(result, []int{nums[0]}, []int{})

	for i := 1; i < len(nums); i++ {
		temp := result[0]
		for j := 0; j < len(result); j++ {
			result[j] = append(result[j], nums[i])
		}
		for k := 0; k < i; k++ {
			result = append(result, []int{nums[k]})
		}
		result = append(result, []int{}, temp)
		fmt.Println(result)
	}

	return result
}

//暴力循环
//func subsets(nums []int)[][]int{
//
//}
