package main

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//使用快慢指针，快指针指向最后的时候，前面的指针刚好是需要删除的位置。
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	result := &ListNode{}
	result.Next = head
	l, r := result, result //先使用一个临时变量指向这个链表。防止第一个元素就是需要删除的。，

	for n > 0 { //先让快指针走 n 次
		r = r.Next
		n--
	}

	for r.Next != nil { //当快指针到最后的时候，说明满指针刚好是需要删除的位置的上一个
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next

	return result.Next
}
