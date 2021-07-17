package main

import "fmt"

//三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。
//
//示例1:
//
// 输入：n = 3
// 输出：4
// 说明: 有四种走法
//示例2:
//
// 输入：n = 5
// 输出：13
//提示:
//
//n范围在[1, 1000000]之间
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/three-steps-problem-lcci

func main() {
	fmt.Println(waysToStep(76))
}

//动态规划
func waysToStep(n int) int {

	var dp []int
	dp = append(dp, 0)
	dp = append(dp, 1) //1 个台阶有 1 种方案，即一次上一步
	dp = append(dp, 2) //2 个台阶有两种方案，即11，2两种方案
	dp = append(dp, 4) //3 个台阶有四种方案，111，12，21，3，
	if n < 4 && n > 0 {
		return dp[n]
	}
	//dp[4] = 7 //七种方案1111,121,211,31,112,121,13 根据规律可知，当前的都为
	i := 4
	for i <= n {
		dp = append(dp, (dp[i-1]+dp[i-2]+dp[i-3])%(1e9+7))
		i++
	}
	return dp[n]
}
