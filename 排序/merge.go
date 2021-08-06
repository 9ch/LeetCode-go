package main

//归并排序

import "fmt"

func main() {
	arr1 := []int{7, 8, 2, 1, 3, 4, 6, 5}
	fmt.Println(mergeSort(arr1))
}

func mergeSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	mid := length >> 1
	left := nums[:mid]
	right := nums[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var result []int

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}

	return result
}
