package main

import (
	"fmt"
	"time"
)

/**

定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	start := time.Now()
	fmt.Println(reverseList(list))
	fmt.Println(time.Since(start))
}

//递归
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}


