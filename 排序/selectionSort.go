package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := random(66)
	fmt.Println(arr)
	s := selectSort(arr)
	fmt.Println(s)
}

//选择排序
//首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
//以此类推，直到所有元素均排序完毕。
func selectSort(arr []int) []int {

	var length = len(arr)

	for i := 0; i < length; i++ {
		var temp = i                      //记录当前的索引位置
		for j := i + 1; j < length; j++ { //记录最小的值的索引
			if arr[temp] > arr[j] {
				temp = j
			}
		}
		arr[i], arr[temp] = arr[temp], arr[i] //交换当前索引位置与最小值索引位置的值
	}
	return arr
}

func random(num int) (result []int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		result = append(result, rand.Intn(num))
	}
	return
}
