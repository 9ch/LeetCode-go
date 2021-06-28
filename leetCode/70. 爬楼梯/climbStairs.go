package main

import (
	"fmt"
	"time"
)

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
*/
func main() {
	n := 44

	start := time.Now()
	fmt.Println(climbStairs(n))
	fmt.Println(time.Since(start))
}

//使用递归
//拆解最小因子。
//LeetCode运行不通过，超过最大时间限制
func climbStairs(n int) int {

	result := 0
	var sum func(n int) int

	sum = func(n int) int {
		if n == 1 { //因为1个的时候只有一种。
			return 1
		}
		if n == 2 { // 两个的时候只有两种。
			return 2
		}
		return sum(n-1) + sum(n-2)
	}
	result = sum(n)

	return result
}

//第二种写法，一样的理念，不过不是重复计算。
func climbStairs2(n int) int {

	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ { //因为值存储了，所以每次计算不需要重复计算，所以时间上非常快。
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
