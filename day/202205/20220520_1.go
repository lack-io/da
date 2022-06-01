package _02205

import "sort"

func findRightInterval(intervals [][]int) []int {
	starts, ans := make([][]int, 0), make([]int, len(intervals))
	for i := range intervals {
		starts = append(starts, []int{intervals[i][0], i})
	}
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})
	starts = append(starts, []int{1e7, -1})
	for i := 0; i < len(starts)-1; i++ {
		cur := sort.Search(len(starts), func(j int) bool {
			return starts[j][0] >= intervals[starts[i][1]][1]
		})
		ans[starts[i][1]] = starts[cur][1]
	}

	return ans
}
