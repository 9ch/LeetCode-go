package main

import "strconv"

//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//输入:
//
//   1
// /   \
//2     3
// \
//  5
//
//输出: ["1->2->5", "1->3"]
//
//解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-paths

func main() {

}

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	if root == nil {
		return result
	}
	var dfs func(node *TreeNode, path string)

	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
			return
		}
		dfs(node.Left, path+"->")
		dfs(node.Right, path+"->")
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
