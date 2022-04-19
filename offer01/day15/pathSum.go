package day15

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	out := make([][]int, 0)
	dfs(root, target, &out, make([]int, 0))
	return out
}

func dfs(root *TreeNode, target int, out *[][]int, cur []int) {
	if root == nil {
		return
	}

	if root.Val == target && (root.Left == nil && root.Right == nil) {
		*out = append(*out, append([]int(nil), append(cur, root.Val)...))
		return
	}

	cur = append(cur, root.Val)
	dfs(root.Left, target-root.Val, out, cur)
	dfs(root.Right, target-root.Val, out, cur)
}
