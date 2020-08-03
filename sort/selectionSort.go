package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := randDom(66)
	s := selectSort(arr)
	fmt.Println(s)
}

//选择排序

func selectSort(arr []int) []int {

	length := len(arr)

	for i := 0; i < length; i++ {
		part := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[part] {
				part = j
			}
		}

		arr[i], arr[part] = arr[part], arr[i]

	}

	return arr
}

func randDom(num int) (result []int) {
	for i := 0; i < num; i++ {
		result = append(result, rand.Intn(num))
	}
	return
}
