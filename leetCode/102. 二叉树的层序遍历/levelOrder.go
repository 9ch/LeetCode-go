package main

import (
	"container/list"
)

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
通过次数172,459提交次数273,043

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
*/

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	return dfs(root, 0, [][]int{})
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {

	if root == nil {
		return res
	}

	if len(res) == level {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}

	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)
	return res
}

//广度优先BFS，需要使用队列。
func bfs(root *TreeNode) [][]int {

	var result [][]int
	if root == nil {
		return result
	}
	//双端队列
	queue := list.New()
	queue.PushFront(root)

	for queue.Len() > 0 {
		var current []int
		listLength := queue.Len()
		for i := 0; i < listLength; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
			current = append(current, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		result = append(result, current)
	}
	return result

}
