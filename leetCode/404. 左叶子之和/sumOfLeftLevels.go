package main

//计算给定二叉树的所有左叶子之和。
//
//示例：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-left-leaves

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//递归
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	midValue := 0

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		midValue = root.Left.Val
	}

	return midValue + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
