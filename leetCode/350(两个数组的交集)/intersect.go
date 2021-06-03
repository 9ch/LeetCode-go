package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 2, 4, 6, 7, 2}
	arr2 := []int{2, 5, 4, 2, 1}

	result := intersect(arr1, arr2)

	fmt.Println(result)

	sort.Ints(arr1)
	sort.Ints(arr2)

	fmt.Println(arr1, arr2)
	fmt.Println(interSect(arr1, arr2))
}

//最简单的实现，通过map映射的方式。时间复杂度，O(n)
func intersect(arr1, arr2 []int) (result []int) {

	m := make(map[int]int)

	for _, v := range arr1 {
		m[v] += 1
	}

	for _, v := range arr2 {
		if m[v] > 0 {
			m[v]--
			result = append(result, v)
		}
	}

	return
}

//有序数组的优化写法
func interSect(arr1, arr2 []int) []int {
	len1, len2 := len(arr1), len(arr2)
	i, j := 0, 0
	var result []int
	if i < len1 && j < len2 {
		if arr1[i] == arr2[j] {
			result = append(result, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}
