package main

import "fmt"

// 最小堆

// 上浮操作
func upAdjust(arr []int) {
	childIndex := len(arr) - 1
	parentIndex := (childIndex - 1) / 2

	temp := arr[childIndex] // 临时存储孩子节点的值
	for childIndex > 0 && temp < arr[parentIndex] {
		arr[childIndex] = arr[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	arr[childIndex] = temp
}

// 下沉操作

func downJust(arr []int, parentIndex int, length int) {
	childIndex := parentIndex*2 + 1 //左节点
	temp := arr[parentIndex]
	for childIndex < length {
		//如果存在右节点，并且右节点的值比左节点的值小
		for childIndex+1 < length && arr[childIndex] > arr[childIndex+1] {
			childIndex++
		}
		//如果父节点的值小于等于孩子节点的值，则直接跳出
		if temp <= arr[childIndex] {
			break
		}

		arr[parentIndex] = arr[childIndex]
		parentIndex = childIndex
		childIndex = childIndex*2 + 1
	}
	arr[parentIndex] = temp
}

// 构建操作

func buildHeap(arr []int) {
	for i := (len(arr) - 2) >> 1; i >= 0; i-- {
		downJust(arr, i, len(arr))
	}
}

func main() {
	arr := []int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0}
	upAdjust(arr)
	fmt.Println(arr)
	arr1 := []int{7, 1, 3, 10, 5, 2, 8, 9, 6}
	buildHeap(arr1)
	fmt.Println(arr1)
}
