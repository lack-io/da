package _02206

import "math"

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		tmp := make([]*TreeNode, 0)
		maxV := math.MinInt
		for len(stack) > 0 {
			node := stack[0]
			if node.Val > maxV {
				maxV = node.Val
			}
			stack = stack[1:]
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		res = append(res, maxV)

		stack = tmp
	}

	return res
}
