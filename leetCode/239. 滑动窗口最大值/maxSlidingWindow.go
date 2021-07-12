package main

import (
	"fmt"
	"sort"
)

/**
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。



进阶：

你能在线性时间复杂度内解决此题吗？



示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
*/
func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}

//滑动窗口
func maxSlidingWindow(nums []int, k int) []int {

	length := len(nums)

	//定义滑动窗口的大小
	var result, window []int //定义返回结果的数组

	temp := 0
	for temp < k { //将前 K 个元素放入到数组中，并去除最大值
		window = append(window, nums[temp])
		temp++
	}
	//将第一次的窗口的最大值求出，并放入到切片中
	result = append(result, MaxValue(window))

	for k < length { //定义循环次数

		window = append(window, nums[k])          //将右边的元素依次放入窗口内
		window = window[1:]                       //将窗口前面的元素删除
		result = append(result, MaxValue(window)) //找出当前窗口的最大值
		k++
	}

	return result
}

//取出当前切片的最大值
func MaxValue(nums []int) int {
	temp := nums
	sort.Ints(temp)
	return nums[len(temp)-1]
}

//取最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
