package main

/**
给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :







返回其后序遍历: [5,6,3,2,4,1].



说明: 递归法很简单，你可以使用迭代法完成此题吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal
*/

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var result []int
	postor(root, result)
	return result
}

func postor(root *Node, result []int) {

	if root != nil {
		for _, v := range root.Children {
			if v != nil {
				postor(v, result)
			}
		}
		result = append(result, root.Val)
	}

}
