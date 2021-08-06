package main

import (
	"fmt"
)

//幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
// 输入： nums = [1,2,3]
// 输出：
//[
//  [3],
// [1],
// [2],
// [1,2,3],
// [1,3],
// [2,3],
// [1,2],
// []
//]
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/power-set-lcci

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(arr1))
	fmt.Println(subsets(arr2))

}

//找规律
func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	if len(nums) < 1 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		temp := make([][]int, len(result))
		copy(temp, result)
		for j := 0; j < len(temp); j++ {
			temp[j] = append(temp[j], nums[i])
		}
		result = append(result, temp...)
	}
	return result
}
