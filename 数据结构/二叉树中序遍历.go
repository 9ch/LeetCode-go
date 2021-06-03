package main

func inorderTraversal(root *TreeNode) []int {
	var result []int

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r != nil {
			dfs(r.Left)
			result = append(result, r.Val)
			dfs(r.Right)
		}
	}
	dfs(root)
	return result
}
