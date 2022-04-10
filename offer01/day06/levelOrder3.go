package day06

func levelOrder3(root *TreeNode) [][]int {
	out := make([][]int, 0)

	if root == nil {
		return out
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	flag := 1
	for len(stack) > 0 {

		length := len(stack)
		q := make([]*TreeNode, 0)
		row := make([]int, 0)

		for i := 0; i < length; i++ {
			node := stack[i]

			switch flag {
			case 1:
				row = append(row, node.Val)
			case -1:
				row = append([]int{node.Val}, row...)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		flag = -1 * flag
		out = append(out, row)
		stack = q
	}

	return out
}
