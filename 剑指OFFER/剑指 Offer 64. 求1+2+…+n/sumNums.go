package main

import (
	"fmt"
)

func main() {
	n1 := 3
	n2 := 9
	fmt.Println(n1, n2)
	fmt.Println(sumNums(n1),sumNums(n2))
}

func sumNums(n int) int {
	ans := 0
	var sum func(n int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
