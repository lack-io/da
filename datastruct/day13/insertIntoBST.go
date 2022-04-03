package day13

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	return insertIntoBST(root, val)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}

	var insert func(root *TreeNode, val int)
	insert = func(node *TreeNode, val int) {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				return
			}
			insert(node.Left, val)
			return
		}
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
			return
		}
		insert(node.Right, val)
	}

	insert(root, val)
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	in := &TreeNode{Val: val}

	if root == nil {
		return in
	}

	node := root
	for node != nil {
		if node.Val > val {
			if node.Left == nil {
				node.Left = in
				break
			}
			node = node.Left
			continue
		}

		if node.Right == nil {
			node.Right = in
			break
		}
		node = node.Right
	}

	return root
}
