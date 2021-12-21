package main

import "fmt"

// 实现最小堆

func main() {
	s := []int{2, 3, 1, 5, 4, 7, 10}
	fmt.Println(s)
	buildHeap(s)
	fmt.Println(s)
}

// 上浮操作
func upJust(arr []int) {
	childIndex := len(arr) - 1 // 最后一个节点
	parentIndex := (childIndex - 1) / 2
	temp := arr[childIndex]
	for parentIndex > 0 && temp < arr[parentIndex] {
		arr[parentIndex] = arr[childIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	arr[childIndex] = temp
}

// 下沉操作
func downJust(arr []int, parentIndex, length int) {
	childIndex := parentIndex*2 + 1 //左孩子节点的位置
	temp := arr[parentIndex]
	for childIndex < length {
		if childIndex+1 < length && arr[childIndex] > arr[childIndex+1] { //如果做孩子节点比右孩子节点小的话，那么就上浮右孩子节点
			childIndex++
		}
		if temp < arr[childIndex] {
			break
		}
		arr[parentIndex] = arr[childIndex]
		parentIndex = childIndex
		childIndex = parentIndex*2 + 1

	}
	arr[parentIndex] = temp
}

func buildHeap(arr []int) {
	for i := (len(arr) - 2) >> 1; i >= 0; i-- {
		downJust(arr, i, len(arr))
	}
}
