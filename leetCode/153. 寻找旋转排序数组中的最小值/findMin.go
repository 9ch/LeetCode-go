package main

import "fmt"

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
*/

func main() {
	a := []int{3, 4, 5, 1, 2}
	b := []int{4, 5, 6, 7, 0, 1, 2}
	c := []int{3, 3, 3, 3, 1}
	//fmt.Println(findMin(a), findMin(b), findMin(c))
	//fmt.Println(findMin2(a), findMin2(b), findMin2(c))
	fmt.Println(findMin3(a), findMin3(b), findMin3(c))
}

//方法1，因为整个数组是升序之后旋转的，所以可以使用增1法来判断，比较简单，O(n)复杂度
func findMin(nums []int) int {

	if len(nums) <= 1 {
		return nums[0]
	}
	result := nums[0]
	for i := 0; i < len(nums); i++ {
		next := i + 1
		if i+1 == len(nums) { //处理超过索引长度，如果超过最大索引，则将下一个索引定位为0
			next = 0
		}
		if nums[next]-nums[i] < 0 { //或者使用 nums[k+1]<v
			return nums[next]
		}
	}
	return result
}

//方法2，暴力循环
func findMin2(nums []int) int {

	min := nums[0]

	for _, v := range nums {
		if v < min { //遍历找出最小值。复杂度O(n)
			min = v
		}
	}
	return min
}

//二分查找
func findMin3(nums []int) int {

	left := 0
	right := len(nums) - 1

	//if len(nums) <= 1 {
	//	return nums[0]
	//}
	//
	//if nums[right] > nums[0] {
	//	return nums[0]
	//}

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}

	return nums[left]
}
