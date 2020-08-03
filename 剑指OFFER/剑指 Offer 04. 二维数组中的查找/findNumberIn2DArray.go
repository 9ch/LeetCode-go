package main

import "fmt"

/**

在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。



示例:

现有矩阵 matrix 如下：

{
  {1,   4,  7, 11, 15},
  {2,   5,  8, 12, 19},
  {3,   6,  9, 16, 22},
  {10, 13, 14, 17, 24},
  {18, 21, 23, 26, 30}
}
给定 target = 5，返回 true。

给定 target = 20，返回 false。



限制：

0 <= n <= 1000

0 <= m <= 1000
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
*/

func main() {
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	//fmt.Println(m)
	fmt.Println(findNumberIn2DArray2(m,500))
}

//暴力解法，直接循环
func findNumberIn2DArray(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

//无视有序特性，放入map里面来查看是否存在。
func findNumberIn2DArray1(matrix [][]int, target int) bool {

	if len(matrix) < 1 {
		return false
	}
	temp := make([]int, len(matrix)*len(matrix[0]))
	k := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			temp[k] = matrix[i][j]
			k++
		}
	}
	for _, v := range temp {
		if v == target {
			return true
		}
	}
	return false
}

//二分法,通过二分法查找没一排的数据。
func findNumberIn2DArray2(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		if findNum(matrix[i], target) {
			return true
		}
	}
	return false
}

//
////回溯法，递归
//func findNumberIn2DArray3(matrix [][]int, target int) bool {
//
//}

//利用二分法查看是否存在
func findNum(nums []int, target int) bool {

	length := len(nums)
	if length < 1 {
		return false

	}
	if length == 1 && nums[0] == target {
		return true
	}

	left := 0
	right := length - 1
	for i := 0; i <= len(nums)/2; i++ {
		mid := (left + right) >> 1
		if target == nums[mid] {
			return true
		}
		if target > nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	return false
}
