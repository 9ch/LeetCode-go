package main

//编写一个函数，检查输入的链表是否是回文的。
//
//
//
//示例 1：
//
//输入： 1->2
//输出： false
//示例 2：
//
//输入： 1->2->2->1
//输出： true
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindrome-linked-list-lcci

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//利用栈，需要两次遍历，O（N）
func isPalindrome(head *ListNode) bool {

	var stack []int //定义一个栈
	var temp = head

	//将元素一次压入栈
	for temp != nil {
		stack = append(stack, temp.Val)
		temp = temp.Next
	}
	//将栈内的元素取出依次与链表比较，如果相等则说明是回文联表
	for i := len(stack) - 1; i > 0; i-- {
		if stack[i] != head.Val {
			return false
		}
		head = head.Next
	}
	return true
}

//快慢指针
//func isPalindrome2(head *ListNode) bool{
//
//}
