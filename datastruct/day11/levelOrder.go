package day11

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder1(root *TreeNode) [][]int {
	out := make([][]int, 0)

	var order func(root *TreeNode, depth int, out *[][]int)
	order = func(root *TreeNode, depth int, out *[][]int) {
		if root == nil {
			return
		}

		if len(*out) == depth {
			*out = append(*out, []int{root.Val})
		} else {
			(*out)[depth] = append((*out)[depth], root.Val)
		}

		if root.Left != nil {
			order(root.Left, depth+1, out)
		}
		if root.Right != nil {
			order(root.Right, depth+1, out)
		}
	}

	order(root, 0, &out)
	return out
}

func LevelOrder2(root *TreeNode) [][]int {
	out := [][]int{}
	if root == nil {
		return out
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		out = append(out, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			out[i] = append(out[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return out
}
