package main

import "fmt"

//某股票每日价格数据包括开盘价、最高价、最低价、收盘价，要求编写一个函数，通过输入一个多日的价格数据，输出每日收盘价的前七日平均价（当日及前六日）
//例如输入：[[100,110,98,104],[104,110,100,108]]输出：
//[[100,110,98,104，104],[104,110,100,108,106]]

func main() {
	s := [][]int{
		{100, 110, 98, 104},
		{104, 110, 100, 108},
		{104, 110, 100, 108},
		{104, 110, 100, 108},
		{104, 110, 100, 108},
		{104, 110, 100, 108},
		{104, 110, 100, 108},
		{104, 110, 100, 108},
	}
	fmt.Println(avg(s))
}

func avg(nums [][]int) [][]int {
	var s []int
	j := 1
	for i := 0; i < len(nums); i++ {
		if j <= 7 {
			s = append(s, nums[i][3])
			a := a(s)
			nums[i] = append(nums[i], a)
			j++
		} else {
			s = append(s, nums[i][3])
			s = s[1:]
			a := a(s)
			nums[i] = append(nums[i], a)
		}
	}

	return nums
}

func a(num []int) int {
	count := 0
	for _, v := range num {
		count += v
	}
	return count / len(num)
}
