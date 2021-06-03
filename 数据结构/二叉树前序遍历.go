package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}

func Pre(root *TreeNode) []int {
	var result []int

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r != nil {
			result = append(result, r.Val)
			dfs(r.Left)
			dfs(r.Right)
		}
	}
	dfs(root)
	return result
}
