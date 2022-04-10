package day06

func levelOrder2(root *TreeNode) [][]int {
	out := make([][]int, 0)

	if root == nil {
		return out
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {

		length := len(stack)
		q := make([]*TreeNode, 0)
		row := make([]int, 0)
		for i := 0; i < length; i++ {
			node := stack[i]

			row = append(row, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		out = append(out, row)
		stack = q
	}

	return out
}
