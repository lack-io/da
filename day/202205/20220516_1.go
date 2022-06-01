package _02205

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)

	flag := false
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if flag {
			return node
		}
		if node.Val == p.Val {
			flag = true
		}
		node = node.Right
	}
	return nil
}

func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < p.Val {
		return inorderSuccessor1(root.Right, p)
	}
	node := inorderSuccessor1(root.Left, p)
	if node == nil {
		return node
	}
	return root
}
