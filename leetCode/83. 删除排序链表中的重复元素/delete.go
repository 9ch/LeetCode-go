package main

import "fmt"

//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
//
//返回同样按升序排列的结果链表。
//
//
//
//示例 1：
//
//
//输入：head = [1,1,2]
//输出：[1,2]
//示例 2：
//
//
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
//
//
//提示：
//
//链表中节点数目在范围 [0, 300] 内
//-100 <= Node.val <= 100
//题目数据保证链表已经按升序排列
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list

func main() {
	//TODO 删除有序链表
	fmt.Println(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针遍历
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var sentinel ListNode
	sentinel.Next = head
	l, r := head, head.Next
	for r != nil {
		if r.Val == l.Val {
			l.Next = r.Next
			r = r.Next
		} else {
			l = r
			r = r.Next
		}
	}
	return sentinel.Next
}

//哈希之类的，有序可以用，但没必要。
