package main

import "fmt"

//给定两个用链表表示的整数，每个节点包含一个数位。
//
//这些数位是反向存放的，也就是个位排在链表首部。
//
//编写函数对这两个整数求和，并用链表形式返回结果。
//
//
//
//示例：
//
//输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
//
//示例：
//
//输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-lists-lcci

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	fmt.Println(addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//主要是处理进位的问题
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	cur1, cur2 := l1, l2 //两个指针
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := result //哨兵节点
	count := 0    //进位
	for cur1 != nil && cur2 != nil {
		val := (cur1.Val + cur2.Val + count) % 10
		count = (cur1.Val + cur2.Val + count) / 10
		cur.Val = val
		cur.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		cur = cur.Next
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur1 == nil {
		for cur2 != nil {
			val := (cur2.Val + count) % 10
			count = (cur2.Val + count) / 10
			cur.Val = val
			cur.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			cur = cur.Next
			cur2 = cur2.Next
		}
	}
	if cur2 == nil {
		for cur1 != nil {
			val := (cur1.Val + count) % 10
			count = (cur1.Val + count) / 10
			cur.Val = val
			cur.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			cur = cur.Next
			cur2 = cur1.Next
		}
	}
	if count == 1 {
		cur.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return result
}
