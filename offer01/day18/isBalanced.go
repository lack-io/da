package day18

func isBalanced(root *TreeNode) bool {
	var recur func(root *TreeNode) int
	recur = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := recur(root.Left)
		if left == -1 {
			return -1
		}
		right := recur(root.Right)
		if right == -1 {
			return -1
		}
		if abs(left-right) < 2 {
			return Max(left, right) + 1
		}
		return -1
	}

	return recur(root) != -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Max(m, n int) int {
	if m < n {
		return n
	}
	return m
}
