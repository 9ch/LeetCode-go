package main

import "fmt"

//给定一个头结点为 head 的非空单链表，返回链表的中间结点。
//
//如果有两个中间结点，则返回第二个中间结点。
//
//
//
//示例 1：
//
//输入：[1,2,3,4,5]
//输出：此列表中的结点 3 (序列化形式：[3,4,5])
//返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
//注意，我们返回了一个 ListNode 类型的对象 ans，这样：
//ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.
//示例 2：
//
//输入：[1,2,3,4,5,6]
//输出：此列表中的结点 4 (序列化形式：[4,5,6])
//由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。
//
//
//提示：
//
//给定链表的结点数介于 1 和 100 之间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/middle-of-the-linked-list

func main() {
	list := &ListNode{
		Vale: 1,
		Next: &ListNode{
			Vale: 2,
			Next: &ListNode{
				Vale: 3,
				Next: &ListNode{
					Vale: 4,
					Next: &ListNode{
						Vale: 5,
						Next: &ListNode{
							Vale: 6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	fmt.Println(middleNode(list))
}

type ListNode struct {
	Vale int
	Next *ListNode
}

// 使用数组，存储链表的每个元素，然后取数字的长度截取
func middleNode(head *ListNode) *ListNode {
	var s []*ListNode

	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	fmt.Println(s)
	return s[len(s)>>1]
}

//还可使用栈

//快慢指针
func middleNode2(head *ListNode) *ListNode {
	pre, cur := head, head
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		pre = pre.Next.Next
		cur = cur.Next
	}
	if pre.Next != nil {
		return cur.Next
	}
	return cur
}
