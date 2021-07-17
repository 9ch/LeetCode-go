package main

import "fmt"

//有一串数字，包含非零和0，把零都放到右边，数字有序的放在左边 如：103002305----> 132350000 要求时间O(n)，空间O(1)
func main() {
	fmt.Println(zero(103002305))
}

func zero(num int) int {
	var temp []int //记录非 0 的数字，并将它们放置数组中，使用的时候反向遍历取出
	count := 0     //记录为 0 的个数
	result := 0
	for num != 0 {
		if t := num % 10; t != 0 {
			temp = append(temp, t)
		} else {
			count++
		}
		num /= 10
	}
	for i := len(temp) - 1; i >= 0; i-- {
		result = result*10 + temp[i]
	}
	for count > 0 {
		result *= 10
		count--
	}
	return result
}
