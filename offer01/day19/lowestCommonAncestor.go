package day19

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	nodes := make(map[*TreeNode]int, 0)
	nodes[root] = 1

	node := root
	queue := make([]*TreeNode, 0)
	queue = append(queue, node)

	level := 0
	for len(queue) > 0 {

		length := len(queue)
		level += 1
		tmp := make([]*TreeNode, 0)
		for i := 0; i < length; i++ {
			node = queue[i]
			if node != root {
				if contains(node, q) && contains(node, p) {
					nodes[node] = level
				}
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}

	out := root
	max := 1
	for k, v := range nodes {
		if v > max {
			max = v
			out = k
		}
	}
	return out
}

func contains(root, node *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == node.Val {
		return true
	}
	return contains(root.Left, node) || contains(root.Right, node)
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
