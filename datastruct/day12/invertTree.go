package day12

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	out := make([]int, 0)

	var inorder func(*TreeNode, *[]int)
	inorder = func(root *TreeNode, out *[]int) {
		if root == nil {
			return
		}
		inorder(root.Left, out)
		*out = append(*out, root.Val)
		inorder(root.Right, out)
	}
	inorder(t, &out)

	return fmt.Sprintf("%v", out)
}

func invertTree1(root *TreeNode) *TreeNode {
	return invertTree(root)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]

			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return root
}
