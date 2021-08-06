package main

import "fmt"

//插入排序
//实现思路：
// 1.从数组的第二个数据开始往前比较，即一开始用第二个数和他前面的一个比较，如果 符合条件（比前面的大或者小，自定义），则让他们交换位置。

// 2.然后再用第三个数和第二个比较，符合则交换，但是此处还得继续往前比较，比如有 5个数8，15，20，45, 17,17比45小，需要交换，但是17也比20小，也要交换，当不需 要和15交换以后，说明也不需要和15前面的数据比较了，肯定不需要交换，因为前 面的数据都是有序的。

// 3.重复步骤二，一直到数据全都排完。
func main() {
	arr := []int{7, 8, 2, 1, 3, 4, 6, 5}
	fmt.Println(arr)
	fmt.Println(insection(arr))
}

//插入排序
func insection(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		pre := i - 1
		cur := nums[i]
		for pre >= 0 && nums[pre] > cur {
			nums[pre+1] = nums[pre]
			pre -= 1
		}
		nums[pre+1] = cur
	}
	return nums
}
