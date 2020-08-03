package main

import (
	"fmt"
	"sort"
)

func main() {

	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{1, 2, 3, 4, 5}

	fmt.Println(maxProfit(prices1))
	fmt.Println(maxProfit(prices2))
	fmt.Println(sort.IntsAreSorted(prices2))
}

func maxProfit(prices []int) int {

	result := 0
	//如果数组长度小于1，或者为倒序排列，则直接返回0
	if len(prices) <= 1 {
		return result
	}

	length := len(prices)

	//如果后一个值比当前值大的话，则把当前值记为一次买入，然后一次进行，即可求出
	for i := 0; i < length; i++ {
		if (i+1) <= length-1 && prices[i] < prices[i+1] {
			result += prices[i+1] - prices[i]
		}
	}

	return result
}
