package _02205

import "math"

func minStickers(stickers []string, target string) int {
	dict := make(map[rune]int, 0)
	for _, c := range target {
		dict[c] += 1
	}

	count := math.MaxInt
	var dfs func([]string, map[rune]int, int)
	dfs = func(ss []string, m map[rune]int, n int) {
		if len(m) == 0 {
			count = min(count, n)
			return
		}
		if len(ss) == 0 && len(m) > 0 {
			count = -1
			return
		}

		text := ss[0]

		flag := false
		for _, c := range text {
			if _, ok := m[c]; ok {
				flag = true
				m[c] -= 1
				if m[c] == 0 {
					delete(m, c)
				}
			}
		}
		if flag {
			dfs(ss, m, n+1)
		} else {
			dfs(ss[1:], m, n)
		}
	}

	for i := 0; i < len(stickers); i++ {
		ss := append([]string{stickers[i]}, stickers[0:i]...)
		ss = append(ss, stickers[i+1:]...)
		dfs(ss, dict, 0)
	}

	return count
}
