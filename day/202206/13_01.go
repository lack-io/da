package _02206

import "sort"

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	res := 0
	for i, v := range sorted {
		if heights[i] != v {
			res += 1
		}
	}
	return res
}
