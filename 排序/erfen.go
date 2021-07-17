package main

import "fmt"

func main() {
	arr1 := []int{1, 3, 5, 6, 8, 10, 22, 89}
	fmt.Println(binarySearch2(arr1, 8))
}

//迭代
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1 //取中间值
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//递归
func binarySearch2(nums []int, target int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	mid := (l + r) >> 1
	if target > nums[mid] {
		return binarySearch2(nums[mid+1:], target)
	} else if target < nums[mid] {
		return binarySearch2(nums[:mid], target)
	} else {
		return mid
	}

}
