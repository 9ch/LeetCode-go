package main

import "fmt"

//统计一个数字在排序数组中出现的次数。
//
//
//
//示例 1:
//
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: 2
//示例 2:
//
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: 0
//
//
//限制：
//
//0 <= 数组长度 <= 50000
//

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(arr)
}

//直接循环查找
func search(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		if target == v {
			count++
		}
	}
	return count
}
