package day6

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[int32]int, 0)
	for _, c := range s {
		m[c] += 1
	}

	for _, c := range t {
		if _, ok := m[c]; !ok {
			return false
		} else {
			m[c] -= 1
			if m[c] == 0 {
				delete(m, c)
			}
		}
	}

	return len(m) == 0
}
