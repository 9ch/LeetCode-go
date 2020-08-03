package main

import "fmt"

/**
给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

你可以返回满足此条件的任何数组作为答案。

 

示例：

输入：[3,1,2,4]
输出：[2,4,3,1]
输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。
 

提示：

1 <= A.length <= 5000
0 <= A[i] <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-array-by-parity
*/
func main() {
	A := []int{3, 1, 2, 4}
	//B := []int{0, 1, 2}
	//fmt.Println(A, append(append(B[:0], B[1:]...),0))
	fmt.Println(sortArrayByPaity3(A))
	//fmt.Println(sortArrayByPaity2(B))
}

//最简单的方法，放到两个数组之中，最好再进行合并
func sortArrayByPaity(A []int) []int {
	var ord []int
	var odd []int

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			ord = append(ord, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}

	return append(ord, odd...)
}

//使用单指针的方法
func sortArrayByPaity2(A []int) []int {

	p := 0
	for i := 0; i < len(A)-1; i++ {
		if A[p]%2 != 0 {
			temp := A[p]
			A = append(append(A[:p], A[p+1:]...), temp)

		} else {
			p++
		}
	}
	return A
}

//使用双指针
func sortArrayByPaity3(A []int) []int {
	k := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[i],A[k] = A[k],A[i]
			k++
		}
	}
	return A
}
