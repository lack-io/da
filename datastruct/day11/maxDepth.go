package day11

func maxDepth(root *TreeNode) int {
	depth := 0

	var order func(root *TreeNode, level int)
	order = func(root *TreeNode, level int) {
		if root == nil {
			depth = max(depth, level+1)
			return
		}

		if root.Left != nil {
			order(root.Left, level+1)
		}
		if root.Right != nil {
			order(root.Right, level+1)
		}
	}

	order(root, 0)
	return depth
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		depth += 1
		for j := 0; j < len(q); j++ {
			node := q[j]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return depth
}

func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepth4(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		depth++
	}
	return depth
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
