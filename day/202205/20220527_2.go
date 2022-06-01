package _02205

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	pre := ""
	if length == 0 {
		return pre
	}

	idx := 0

	exit := false
	for {
		if exit {
			break
		}

		c := uint8(0)
		for i := 0; i < length; i++ {
			if idx >= len(strs[i]) {
				exit = true
				break
			}
			if i == 0 {
				c = strs[i][idx]
			}

			if strs[i][idx] != c {
				exit = true
				break
			}
		}
		if !exit {
			idx += 1
			pre += string(c)
		}
	}

	return pre
}
