package day14

func findTarget1(root *TreeNode, k int) bool {
	stack := make([]*TreeNode, 0)

	hm := make(map[int]struct{}, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(hm) == 0 {
			hm[node.Val] = struct{}{}
		} else {
			_, ok := hm[k-node.Val]
			if ok {
				return true
			} else {
				hm[node.Val] = struct{}{}
			}
		}
		node = node.Right
	}

	return true
}
