package main

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3



示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}

//递归，通过判断left.Left.val 和right.Right.val是否相等来计算。
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)

}

func recur(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return recur(left.Left, right.Right) && recur(left.Right, right.Left)
}

//解法2 中序遍历出来之后，再进行判断i 和n-i的索引位置上面的值是否相等。
