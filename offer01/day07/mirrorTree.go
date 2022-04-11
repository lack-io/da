package day07

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		mirrorTree(root.Left)
	}
	if root.Right != nil {
		mirrorTree(root.Right)
	}
	return root
}
