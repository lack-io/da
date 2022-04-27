package day10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal1(root *TreeNode) []int {
	out := make([]int, 0)

	var traversal func(*TreeNode, *[]int)
	traversal = func(root *TreeNode, out *[]int) {
		*out = append(*out, root.Val)
		traversal(root.Left, out)
		traversal(root.Right, out)
	}
	traversal(root, &out)

	return out
}

func PreorderTraversal2(root *TreeNode) []int {
	vals := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return vals
}

func PreorderTraversal3(root *TreeNode) []int {
	vals := make([]int, 0)
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				vals = append(vals, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			vals = append(vals, p1.Val)
		}
		p1 = p1.Right
	}
	return vals
}
