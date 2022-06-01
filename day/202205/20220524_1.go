package _02205

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(root *TreeNode, val int) bool
	dfs = func(root *TreeNode, val int) bool {
		if root == nil {
			return true
		}
		return root.Val == val && dfs(root.Left, val) && dfs(root.Left, val)
	}
	return dfs(root.Left, root.Val) && dfs(root.Left, root.Val)
}
