package _02206

func findBottomLeftValue(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	var target *TreeNode
	for len(stack) > 0 {
		target = stack[len(stack)-1]
		tmp := make([]*TreeNode, 0)
		for len(stack) > 0 {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}

		stack = tmp
	}

	return target.Val
}
