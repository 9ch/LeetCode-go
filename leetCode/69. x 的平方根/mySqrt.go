package main

import "fmt"

/**
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sqrtx
*/
func main() {
	num := 4
	//fmt.Println(num)
	fmt.Println(mySqrt(num))
}

//目前想到的办法就是初中数学学得二分法了。。。
func mySqrt(x int) int {

	if x == 1 {
		return 1
	}
	if x == 0 {
		return 0
	}
	mid := 0
	left := 0
	right := x
	for left <= right { //当左边值小于右边值的时候，进行计算，（核心，如果到临界值的时候，会出现右边值小于左边值的情况，此时，右边值就是我们需要的。）
		mid = (left + right) / 2 //取中间值
		if mid*mid < x {         //当中间值的平方小于变量的话，则将左边值+1
			left = mid + 1
		} else if mid*mid > x { //当中间值的平方大于变量的话，则将右边值-1
			right = mid - 1
		} else {
			return mid
		}
	}

	return right
}

/**
以8为例：
第1次循环： 0 8
第2次循环： 0 3
第3次循环： 1 3
第4次循环： 2 3
第5次循环： 3 3
第6次循环： 3 2
此时的右边值2即为目标值。
*/
