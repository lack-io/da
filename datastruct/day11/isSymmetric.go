package day11

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return compare(root.Left, root.Right)
}

func compare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

func isSymmetric2(root *TreeNode) bool {
	q := make([]*TreeNode, 0)

	u, v := root, root
	q = append(q, u, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]

		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, u.Right)
		q = append(q, v.Right)
		q = append(q, v.Left)
	}

	return true
}
