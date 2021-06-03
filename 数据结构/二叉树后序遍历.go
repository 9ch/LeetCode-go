package main

func main() {

}

func Next(root *TreeNode) []int {
	var result []int

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r != nil {
			dfs(r.Left)
			dfs(r.Right)
			result = append(result, r.Val)
		}
	}
	dfs(root)
	return result
}
