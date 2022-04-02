package day12

func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var dps func(root *TreeNode, sum int) bool
	dps = func(root *TreeNode, sum int) bool {
		if root == nil {
			return sum == targetSum
		}

		sum += root.Val
		if root.Left == nil && root.Right != nil {
			return dps(root.Right, sum)
		}
		if root.Left != nil && root.Right == nil {
			return dps(root.Left, sum)
		}

		return dps(root.Left, sum) || dps(root.Right, sum)
	}

	return dps(root, 0)
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queNode := []*TreeNode{}
	queVal := []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) != 0 {
		now := queNode[0]
		queNode = queNode[1:]
		temp := queVal[0]
		queVal = queVal[1:]
		if now.Left == nil && now.Right == nil {
			if temp == targetSum {
				return true
			}
			continue
		}
		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+temp)
		}
		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+temp)
		}
	}
	return false
}
