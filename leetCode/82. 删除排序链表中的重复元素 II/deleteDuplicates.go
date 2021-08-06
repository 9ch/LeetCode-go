package main

//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
//
//返回同样按升序排列的结果链表。
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//示例 2：
//
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
//
//
//提示：
//
//链表中节点数目在范围 [0, 300] 内
//-100 <= Node.val <= 100
//题目数据保证链表已经按升序排列
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

//递归
func deleteDuplicates2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates2(head.Next)
	} else {
		temp := head.Next
		for temp != nil && head.Val == temp.Val {
			temp = temp.Next
		}
		return deleteDuplicates2(temp)
	}
	return head
}
