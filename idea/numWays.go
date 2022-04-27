package idea

// https://leetcode-cn.com/problems/chuan-di-xin-xi/
func numWays(n int, relation [][]int, k int) int {
	ways := make(map[int]map[int]struct{}, 0)

	for _, rel := range relation {
		v, ok := ways[n]
		if !ok {
			v = make(map[int]struct{})
			ways[n] = v
		}
		v[rel[1]] = struct{}{}
	}

	nums := 0
	dfs(0, n-1, k, ways, &nums)
	return nums
}

func dfs(from, n, k int, ways map[int]map[int]struct{}, nums *int) {
	if k == 0 {
		if from == n {
			*nums += 1
		}
		return
	}
	if way, ok := ways[from]; ok {
		for to, _ := range way {
			dfs(to, n, k-1, ways, nums)
		}
	}
}
