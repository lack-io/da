package day06

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	out := make([]int, 0)

	if root == nil {
		return out
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {

		length := len(stack)
		q := make([]*TreeNode, 0)
		for i := 0; i < length; i++ {
			node := stack[i]

			out = append(out, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		stack = q
	}

	return out
}
