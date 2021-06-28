package main

import "fmt"

//一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
//
//
//
//示例 1:
//
//输入: [0,1,3]
//输出: 2
//示例2:
//
//输入: [0,1,2,3,4,5,6,7,9]
//输出: 8
//
//
//限制：
//
//1 <= 数组长度 <= 10000
//

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	a2 := []int{0}
	fmt.Println(missingNumber(a1))
	fmt.Println(missingNumber(a2))
}

//暴力循环
func missingNumber(nums []int) int {
	for k, v := range nums {
		if k != v {
			return k
		}
	}
	return len(nums) //处理特殊情况，例如【0，1】，【0】
}

//二分法
func missingNumber2(nums []int) int {
	//TODO
}
