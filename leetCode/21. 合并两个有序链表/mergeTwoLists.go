package main

import "fmt"

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	fmt.Println(mergeTwoLists(l1, l2))
}

//迭代，通过设置哨兵节点。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	temp := &ListNode{}
	result := temp //核心点，设置哨兵节点，即头结点，这样修改上面的链接，对哨兵节点来说没影响，只需要到时候返回哨兵节点的Next就行了。
	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			temp.Next = l1 //哪个值小，就将当前节点设置成哪个节点。
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 == nil {
		temp.Next=l2
	}else {
		temp.Next=l1
	}
	return result.Next
}
