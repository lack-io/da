package _02206

func findFrequentTreeSum(root *TreeNode) (ans []int) {
	cnt := map[int]int{}
	maxCnt := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		cnt[sum]++
		if cnt[sum] > maxCnt {
			maxCnt = cnt[sum]
		}
		return sum
	}
	dfs(root)

	for s, c := range cnt {
		if c == maxCnt {
			ans = append(ans, s)
		}
	}
	return
}
