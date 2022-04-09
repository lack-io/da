package day05

import "math"

func firstUniqChar(s string) byte {
	sets := make(map[int32]int)

	for _, c := range s {
		if _, ok := sets[c]; ok {
			sets[c] = -1
		} else {
			sets[c] = 1
		}
	}

	if len(sets) == 0 {
		return ' '
	}

	index := math.MaxInt
	for _, v := range sets {
		if v < index && v != -1 {
			index = v
		}
	}

	return s[index]
}
