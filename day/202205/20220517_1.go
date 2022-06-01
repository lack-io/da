package _02205

func isAlienSorted(words []string, order string) bool {
	length := len(words)
	if length <= 1 {
		return true
	}

	m := map[rune]int{}
	for i, c := range order {
		m[c] = i
	}

	for i := 1; i < length; i++ {
		pLen, qLen := len(words[i-1]), len(words[i])
		p, q := 0, 0
		for p < pLen || q < qLen {
			if p == pLen {
				break
			}

			if q == qLen {
				return false
			}

			if words[i-1][p] != words[i][q] {
				if m[rune(words[i-1][p])] > m[rune(words[i][q])] {
					return false
				}
				break
			}
		}

		p += 1
		q += 1
	}

	return true
}
