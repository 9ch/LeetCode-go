package main

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//根据栈的原理查询
func isPalindRome(head *ListNode) bool {

	var stack []int //定义一个栈，存放节点
	pre := head     //临时的一个指针变量

	//将数值压入到栈中
	for pre != nil {
		stack = append(stack, pre.Val)
		pre = pre.Next
	}

	//依次遍历栈弹出的值是否与当前节点的值相等，如果不想等，这说明不是回文
	for i := len(stack) - 1; i > 0; i-- {
		if stack[i] != head.Val {
			return false
		} else {
			head = head.Next
		}
	}

	return true

}
