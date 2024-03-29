package main

/**
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。




进阶：

你能用 O(1)（即，常量）内存解决此问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

//快慢指针，快指针走两步，满指针走一步，如果快指针追上满指针，则说明有环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	cur := head.Next

	for cur != nil && head != nil && cur.Next != nil {
		if head == cur {
			return true
		}
		cur = cur.Next.Next
		head = head.Next
	}
	return false
}

//通过哈希判断是否有环
func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	m := make(map[*ListNode]struct{})
	cur := head.Next
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true
		}
		m[cur] = struct{}{}
		cur = cur.Next
	}
	return false
}
