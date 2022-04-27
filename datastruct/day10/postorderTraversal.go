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

func PostorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	out := make([]int, 0)
	var mostRight *TreeNode
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
				printEdge(cur.Left, &out)
			}
		}
		cur = cur.Right
	}
	printEdge(root, &out)
	return out
}

func printEdge(node *TreeNode, out *[]int) {
	tail := reverse(node)
	cur := tail
	for cur != nil {
		*out = append(*out, cur.Val)
		cur = cur.Right
	}
	reverse(tail)
}

func reverse(node *TreeNode) *TreeNode {
	var pre, next *TreeNode
	for node != nil {
		next = node.Right
		node.Right = pre
		pre = node
		node = next
	}
	return pre
}
