package main

import "fmt"

/**
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

 

示例 1：

输入：n = 2
输出：1
示例 2：

输入：n = 5
输出：5
 

提示：

0 <= n <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof
*/

func main() {
	fmt.Println(fib2(5))
}

//递归 超出时间限制。
func fib(n int) int {

	if n == 0 {
		return 0
	}
	var sum func(n int) int
	sum = func(n int) int {
		if n == 1 {
			return 1
		}
		return fib(n-1) + fib(n-2)
	}
	result := sum(n) % 1000000007
	return result

}

//动态规划
func fib2(n int) int {

	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %=1000000007
	}
	return dp[n]%1000000007
}
