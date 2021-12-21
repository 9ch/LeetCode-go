package main

import "container/list"

//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
//
//
//示例 1:
//
//
//
//输入: [1,2,3,null,5,null,4]
//输出: [1,3,4]
//示例 2:
//
//输入: [1,null,3]
//输出: [1,3]
//示例 3:
//
//输入: []
//输出: []
//
//
//提示:
//
//二叉树的节点个数的范围是 [0,100]
//-100 <= Node.val <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-right-side-view

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			ele := queue.Remove(queue.Back()).(*TreeNode)
			if ele.Left != nil {
				queue.PushFront(ele.Left)
			}
			if ele.Right != nil {
				queue.PushFront(ele.Right)
			}
			if i == length-1 {
				result = append(result, ele.Val)
			}
		}
	}
	return result
}
