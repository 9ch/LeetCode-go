package main

import (
	"fmt"
	"sort"
)

//一个数对 (a,b) 的 数对和 等于 a + b 。最大数对和 是一个数对数组中最大的 数对和 。

// 比方说，如果我们有数对 (1,5) ，(2,3) 和 (4,4)，最大数对和 为 max(1+5, 2+3, 4+4) = max(6, 5, 8) = 8 。
// 给你一个长度为 偶数 n 的数组 nums ，请你将 nums 中的元素分成 n / 2 个数对，使得：

// nums 中每个元素 恰好 在 一个 数对中，且
// 最大数对和 的值 最小 。
// 请你在最优数对划分的方案下，返回最小的 最大数对和 。

//

// 示例 1：

// 输入：nums = [3,5,2,3]
// 输出：7
// 解释：数组中的元素可以分为数对 (3,3) 和 (5,2) 。
// 最大数对和为 max(3+3, 5+2) = max(6, 7) = 7 。
// 示例 2：

// 输入：nums = [3,5,4,2,4,6]
// 输出：8
// 解释：数组中的元素可以分为数对 (3,5)，(4,4) 和 (6,2) 。
// 最大数对和为 max(3+5, 4+4, 6+2) = max(8, 8, 8) = 8 。
//

// 提示：

// n == nums.length
// 2 <= n <= 105
// n 是 偶数 。
// 1 <= nums[i] <= 105

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimize-maximum-pair-sum-in-array

func main() {
	fmt.Println(minPairSum([]int{3, 5, 2, 3}))
	fmt.Println(minPairSum([]int{3, 5, 4, 2, 4, 6}))
	fmt.Printf("%v,%+v", "a", "a")

}

//让最大的加上最小的，即可满足，可以边计算最小值
func minPairSum(nums []int) int {
	sort.Ints(nums)        //先排序
	i, j := 0, len(nums)-1 //定义两个指针
	var max int            //存储最大值
	for i < j {
		if n := nums[i] + nums[j]; n > max { //边遍历，边计算出最大值，不需要太多额外空间。
			max = n
		}
		i++
		j--
	}
	return max
}
