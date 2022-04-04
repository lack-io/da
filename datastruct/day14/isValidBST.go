package day14

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST1(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)

	cur := math.MinInt
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == math.MinInt || cur < node.Val {
			cur = node.Val
		} else {
			return false
		}
		node = node.Right
	}

	return true
}

func isValidBST2(root *TreeNode) bool {
	return compare(root, math.MinInt64, math.MaxInt64)
}

func compare(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return compare(root.Left, lower, root.Val) && compare(root.Right, root.Val, upper)
}
