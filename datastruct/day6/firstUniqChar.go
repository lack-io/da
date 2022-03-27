package day6

func FirstUniqChar(s string) int {
	sets := make(map[int32]int, 0)
	for _, c := range s {
		if _, ok := sets[c]; ok {
			sets[c] += 1
		} else {
			sets[c] = 1
		}
	}

	for i, c := range s {
		if sets[c] == 1 {
			return i
		}
	}
	return -1
}
