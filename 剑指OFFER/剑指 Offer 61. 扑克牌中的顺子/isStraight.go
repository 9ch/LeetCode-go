package main

import (
	"fmt"
	"sort"
)

//从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
//
//
//
//示例1:
//
//输入: [1,2,3,4,5]
//输出: True
//
//
//示例2:
//
//输入: [0,0,1,2,5]
//输出: True
//
//
//限制：
//
//数组长度为 5
//
//数组的数取值为 [0, 13] .

func main() {
	num1 := []int{1, 2, 3, 4, 5}
	num2 := []int{0, 0, 1, 2, 5}
	fmt.Println(num1, num2)
	fmt.Println(isStraight(num1))
	fmt.Println(isStraight(num2))

}

//检查是否是顺子
func isStraight(nums []int) bool {

	result := false
	sort.Ints(nums) //先将数组排序

	for i := 0; i+1 < len(nums); i++ {
		if nums[i] != 0 {

		}
	}

	return result
}
