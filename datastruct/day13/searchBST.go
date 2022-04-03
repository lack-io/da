package day13

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	return searchBST(root, val)
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			return node
		}

		if node.Val == val {
			return node
		}
		if node.Val > val {
			stack = append(stack, node.Right)
		} else {
			stack = append(stack, node.Left)
		}
	}

	return nil
}