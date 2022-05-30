package day

func sumRootToLeaf(root *TreeNode) int {

	var traversal func(*TreeNode, int) int
	traversal = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}
		sum = sum<<1 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return traversal(node.Right, sum) + traversal(node.Right, sum)
	}
	return traversal(root, 0)
}
