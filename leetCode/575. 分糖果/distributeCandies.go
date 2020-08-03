package main

import (
	"fmt"
	"sort"
)

/**
给定一个偶数长度的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。你需要把这些糖果平均分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。

示例 1:

输入: candies = [1,1,2,2,3,3]
输出: 3
解析: 一共有三种种类的糖果，每一种都有两个。
     最优分配方案：妹妹获得[1,2,3],弟弟也获得[1,2,3]。这样使妹妹获得糖果的种类数最多。
示例 2 :

输入: candies = [1,1,2,3]
输出: 2
解析: 妹妹获得糖果[2,3],弟弟获得糖果[1,1]，妹妹有两种不同的糖果，弟弟只有一种。这样使得妹妹可以获得的糖果种类数最多。
注意:

数组的长度为[2, 10,000]，并且确定为偶数。
数组中数字的大小在范围[-100,000, 100,000]内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/distribute-candies
*/
func main() {
	codies := []int{1, 1, 2, 2, 3, 3}
	codies1 := []int{1, 1, 2, 3}
	codies3 := []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3}

	fmt.Println(codies, codies1)
	fmt.Println(distributeCandies(codies), distributeCandies(codies1), distributeCandies(codies3))
	fmt.Println(distributeCandies2(codies), distributeCandies2(codies1), distributeCandies2(codies3))
}

//增1法，比较相邻的两个值，如果不一致，则说明需要有不同的糖果。将值+1.
func distributeCandies(candies []int) int {

	result := 1        //默认值为1，因为偶数个值，说明最小的也是1个。
	sort.Ints(candies) //将数组进行排序。

	for k, v := range candies {
		if result > len(candies)>>1 { //分配的值最多等于len(candies)/2
			result = len(candies) >> 1
		}
		if k+1 == len(candies) {
			return result
		}
		if v != candies[k+1] {
			result++
		}
	}

	return result
}

//放到map集合里面，因为map的值是唯一的。
func distributeCandies2(candies []int) int {

	result := 0
	maps := make(map[int]int)

	for _, v := range candies {
		maps[v] = 1 //无需在乎值，值为任意数都可以。
	}
	fmt.Println(maps)
	if len(maps) < len(candies)/2 {
		result = len(maps)
	} else {
		result = len(candies) / 2
	}

	return result
}
