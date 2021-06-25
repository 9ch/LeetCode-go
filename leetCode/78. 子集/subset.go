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
	//nums1 := []int{1, 2, 3}
	//nums2 := []int{1, 2}
	//nums3 := []int{1}
	fmt.Println(subsets1(nums))
	//fmt.Println(subsets1(nums1))
	//fmt.Println(subsets1(nums2))
	//fmt.Println(subsets1(nums3))
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
	}

	return result
}

//核心思想是遍历的时候，每个元素都把当前结果集里面的数据复制出来，并把当前当前值放进去，最后再放到结果集里面。
func subsets1(nums []int) [][]int {

	length := len(nums)
	fmt.Println(nums)
	result := [][]int{[]int{}} //定义几个空的结果集

	for i := 0; i < length; i++ { //遍历当前元素
		temp := result //复制一份当前结果集出来
		fmt.Println(temp)
		for j := 0; j < len(temp); j++ { //将当前元素追加到复制出来的结果集的每一个子数组里面。
			temp[j] = append(temp[j], nums[i])
		}
		result = append(result, temp...) //将当前已经添加好的结果集再次追加到先前的结果集里面。
	}
	return result
}
