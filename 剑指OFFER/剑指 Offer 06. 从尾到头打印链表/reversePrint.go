package main

import "fmt"

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

 

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
 

限制：

0 <= 链表长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil},
		},
	}

	fmt.Println(reversePrint(&head))
	fmt.Println(reversePrint2(&head))
}

//递归法
func reversePrint(head *ListNode) []int {

	var result []int

	var df func(h *ListNode)

	df = func(h *ListNode) {
		if h == nil {
			return
		}
		df(h.Next)
		result = append(result, h.Val)
	}
	df(head)

	return result
}

//暴力循环
func reversePrint2(head *ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return result
}
