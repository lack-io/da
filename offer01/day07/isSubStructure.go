package day07

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var compare func(a *TreeNode, b *TreeNode) bool
	compare = func(a *TreeNode, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil && b != nil {
			return false
		}
		if a != nil && b == nil {
			return false
		}
		ok := a.Val == b.Val
		if b.Left != nil {
			ok = ok && compare(a.Left, b.Left)
		}
		if b.Right != nil {
			ok = ok && compare(a.Right, b.Right)
		}
		return ok
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, A)

	for len(stack) > 0 {

		length := len(stack)
		q := make([]*TreeNode, 0)
		for i := 0; i < length; i++ {
			node := stack[i]

			if compare(node, B) {
				return true
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		stack = q
	}

	return false
}
