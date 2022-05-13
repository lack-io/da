package day

func oneEditAway(first string, second string) bool {
	plen := len(first)
	qlen := len(second)

	switch {
	case plen > qlen:
		// "aaabb"
		// "aabb"
		idx := qlen
		for i := 0; i < qlen; i++ {
			if first[i] != second[i] {
				idx = i
				break
			}
		}
		if idx == qlen {
			return plen-idx == 1
		}
		return first[idx+1:] == second[idx:]

	case plen < qlen:

		idx := plen
		for i := 0; i < plen; i++ {
			if first[i] != second[i] {
				idx = i
				break
			}
		}
		if idx == plen {
			return plen-idx == 1
		}
		return first[idx:] == second[idx+1:]

	default:

		idx := plen
		for i := 0; i < plen; i++ {
			if first[i] != second[i] {
				idx = i
				break
			}
		}
		if idx == plen {
			return true
		}
		return first[idx+1:] == second[idx+1:]
	}
}

func oneEditAway1(first string, second string) bool {
	if len(first) == 0 {
		return len(second) <= 1
	}
	if len(second) == 0 {
		return len(first) <= 1
	}
	if first[0] != second[0] {
		if len(first) == len(second) {
			return first[1:] == second[1:]
		} else if len(first) > len(second) {
			return first[1:] == second
		}
		return first == second[1:]
	}
	return oneEditAway(first[1:], second[1:])
}
