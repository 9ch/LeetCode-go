package main

/**
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-node-in-a-bst
*/

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}

	if key < root.Val { //key比当前值小的时候，去左节点查找
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val { //key比当前值大的时候，去有节点查找删除
		root.Right = deleteNode(root.Right, key)
	} else { //当查找到了进行操作
		if root.Left == nil { //如果当前节点的左节点为空的话，直接返回右节点
			return root.Right
		} else if root.Right == nil { //如果当前节点的右节点为空的话，直接返回左节点
			return root.Left
		} else { //当节点都存在的时候，将当前节点的左节点挪到当前节点的右节点的最左节点上面。
			tree := root.Right
			for tree.Left != nil {
				tree = tree.Left
			}
			tree.Left = root.Left
			return root.Right
		}
	}
	return root
}
