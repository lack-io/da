package day15

func kthLargest(root *TreeNode, k int) int {
	res := 0

	stack := make([]*TreeNode, 0)

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k -= 1
		if k == 0 {
			res = node.Val
		}

		node = node.Left
	}

	return res
}
