package main

import (
	"fmt"
	"sort"
)

func main() {

	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{1, 2, 3, 4, 5}
	prices3 := []int{5, 4, 3, 2, 1}

	fmt.Println(maxProfit(prices1))
	fmt.Println(maxProfit(prices2))
	fmt.Println(maxProfit(prices3))
	fmt.Println(sort.IntsAreSorted(prices2))
}

func maxProfit(prices []int) int {

	result := 0

	if len(prices) < 1 {
		return result
	}

	length := len(prices)
	//只有当前面的值比后面的小时，才会盈利
	for i := 0; i < length; i++ {
		if i+1 <= length-1 && prices[i] < prices[i+1] {
			result = result + prices[i+1] - prices[i]
		}
	}
	return result
}
