package day18

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth += 1
	return max(dfs(root.Left, depth), dfs(root.Right, depth))
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	node := root
	queue := make([]*TreeNode, 0)
	queue = append(queue, node)

	depth := 0
	for len(queue) > 0 {
		depth += 1
		length := len(queue)
		tmp := make([]*TreeNode, 0)
		for i := 0; i < length; i++ {
			node = queue[i]

			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}

		queue = tmp
	}

	return depth
}