package main

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表
//小范围替换
func reserveList(list *ListNode) *ListNode {

	var result *ListNode
	pre := list

	for pre != nil {
		temp := pre.Next
		pre.Next = result
		result = pre
		pre = temp
	}

	return result
}
