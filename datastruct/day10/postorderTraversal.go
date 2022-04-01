package day10

func postorderTraversal1(root *TreeNode) []int {
	out := make([]int, 0)

	var traversal func(*TreeNode, *[]int)
	traversal = func(root *TreeNode, out *[]int) {
		if root == nil {
			return
		}
		traversal(root.Left, out)
		traversal(root.Right, out)
		*out = append(*out, root.Val)
	}
	traversal(root, &out)

	return out
}

func postorderTraversal2(root *TreeNode) []int {
	out := make([]int, 0)
	stack := make([]*TreeNode, 0)

	var prev *TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			out = append(out, node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}

	return out
}
