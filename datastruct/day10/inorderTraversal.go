package day10

func InorderTraversal1(root *TreeNode) []int {
	out := make([]int, 0)

	var inorder func(*TreeNode, *[]int)
	inorder = func(root *TreeNode, out *[]int) {
		if root == nil {
			return
		}
		inorder(root.Left, out)
		*out = append(*out, root.Val)
		inorder(root.Right, out)
	}
	inorder(root, &out)

	return out
}

func InorderTraversal2(root *TreeNode) []int {
	out := make([]int, 0)
	stack := make([]*TreeNode, 0)

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, node.Val)
		node = node.Right
	}

	return out
}

func InorderTraversal3(root *TreeNode) []int {
	vals := make([]int, 0)
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			vals = append(vals, p1.Val)
			p2.Right = nil
			p1 = p1.Right
		} else {
			vals = append(vals, p1.Val)
			p1 = p1.Right
		}
	}
	return vals
}
