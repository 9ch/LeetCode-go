package main

//给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//
//
//示例 1：
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：tail connects to node index 1
//解释：链表中有一个环，其尾部连接到第二个节点。
//示例 2：
//
//
//
//输入：head = [1,2], pos = 0
//输出：tail connects to node index 0
//解释：链表中有一个环，其尾部连接到第一个节点。
//示例 3：
//
//
//
//输入：head = [1], pos = -1
//输出：no cycle
//解释：链表中没有环。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/linked-list-cycle-lcci

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//哈希表的方式
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	m[head] = true
	for head.Next != nil {
		if _, ok := m[head.Next]; ok {
			return head.Next
		}
		m[head.Next] = true
		head = head.Next
	}
	return nil
}

//快慢指针
//TODO
