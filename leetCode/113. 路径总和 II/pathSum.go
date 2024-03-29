package main

import "fmt"

/**
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
*/

func main() {
	fmt.Println("hello world")
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int

	var path []int
	var dfs func(node *TreeNode, int2 int)

	dfs = func(node *TreeNode, n int) {
		if node == nil {
			return
		}
		n -= node.Val
		path = append(path, node.Val)
		if n == 0 && node.Left == nil && node.Right == nil {
			result = append(result, append([]int(nil), path...))
			return
		}
		dfs(node.Left, n)
		dfs(node.Right, n)
		path = path[:len(path)-1]
	}
	dfs(root, sum)
	return result
}
