package main

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

 

示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。
 

提示：

1 <= nums.length <= 50000
1 <= nums[i] <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof
*/

func main() {

}

//使用临时变量。
func exchange(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	var left, right []int

	for _, v := range nums {

		if v%2 == 0 {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	return append(left, right...)
}

//不使用临时变量,使用快慢指针
func exchange2(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	k:=0
	for i := 0; i < len(nums);i++ {
		if nums[i] %2 == 1 {
			nums[i],nums[k] = nums[k],nums[i]
			k++
		}
	}
	return nums
}
