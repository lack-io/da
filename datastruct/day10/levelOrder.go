package day10

func levelTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	out := make([]int, 0)

	node := root
	queue := make([]*TreeNode, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]*TreeNode, 0)
		for i := 0; i < length; i++ {
			node = queue[i]

			out = append(out, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}

		queue = tmp
	}

	return out
}
