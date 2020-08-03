package main

import (
	"fmt"
	"sort"
)

/**
给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。

如果不能形成任何面积不为零的三角形，返回 0。

 

示例 1：

输入：[2,1,2]
输出：5
示例 2：

输入：[1,2,1]
输出：0
示例 3：

输入：[3,2,3,4]
输出：10
示例 4：

输入：[3,6,2,3]
输出：8
 

提示：

3 <= A.length <= 10000
1 <= A[i] <= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-perimeter-triangle
*/
func main() {
	a1 := []int{3, 6, 2, 3}
	a2 := []int{3, 2, 3, 4}
	a3 := []int{1, 2, 1}

	fmt.Println(largestPerimeter(a1),largestPerimeter(a2),largestPerimeter(a3))
}

//三角形定理：任意两边之和大于第三边。根据此定理来计算
func largestPerimeter(A []int) int {

	result := 0
	sort.Ints(A) //先将数组排序

	for i := len(A) - 1; i >= 0; i-- { //倒序寻找，找到的第一个即为最大的
		if i <2 {
			return 0
		}
		if A[i] < (A[i-1] + A[i-2]) {
			return A[i] + A[i-1] + A[i-2]
		}
	}

	return result
}
