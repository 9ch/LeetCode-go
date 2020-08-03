package main

/**
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

 

示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

//遍历,双指针（核心）
func getKthFromEnd(head *ListNode, k int) *ListNode {

	pre := head
	next := head

	for i := 0; i < k; i++ { //先让前指针多走K步
		next = next.Next
	}

	for next != nil { //循环遍历，一直到快指针走到链表尾部。
		next,pre = next.Next,pre.Next
	}

	return pre

}
