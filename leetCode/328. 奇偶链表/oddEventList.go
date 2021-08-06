package main

import "fmt"

//给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
//
//请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
//
//示例 1:
//
//输入: 1->2->3->4->5->NULL
//输出: 1->3->5->2->4->NULL
//示例 2:
//
//输入: 2->1->3->5->6->4->7->NULL
//输出: 2->3->6->7->1->5->4->NULL
//说明:
//
//应当保持奇数节点和偶数节点的相对顺序。
//链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/odd-even-linked-list

func main() {
	s := &ListNode{
		Val: 1,
		Next: &ListNode{
			2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}},
		},
	}
	a := oddEvenList(s)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//利用数组，奇偶的都放入对应数组，然后遍历添加即可,LeetCode 不通过，内存泄露 OOM
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var odd []*ListNode
	var even []*ListNode
	count := 0 //判断是否为奇偶
	for head != nil {
		if count%2 == 0 {
			odd = append(odd, head)
		} else {
			even = append(even, head)
		}
		count++
		head = head.Next
	}
	sen := &ListNode{}
	head = sen
	for _, v := range odd {
		head.Next = v
		head = head.Next
	}
	for _, v := range even {
		head.Next = v
		head = head.Next
	}
	return sen.Next
}

func oddEvenList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	evenHead := head.Next
	even := evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
