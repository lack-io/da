package day28

func permutation(s string) []string {
	sets := make(map[string]struct{}, 0)

	result := make([]string, 0)
	var dps func(s, target string, sets map[string]struct{}, result *[]string)
	dps = func(s, target string, sets map[string]struct{}, result *[]string) {
		if s == "" {
			if _, ok := sets[target]; !ok {
				*result = append(*result, target)
				sets[target] = struct{}{}
			}

			return
		}

		for i, c := range s {
			other := ""
			if i == len(s)-1 {
				other = s[:len(s)-1]
			} else {
				other = s[:i] + s[i+1:]
			}

			dps(other, target+string(c), sets, result)
		}
	}

	dps(s, "", sets, &result)

	return result
}

func permutation1(s string) []string {
	result := make([]string, 0)
	c := []byte(s)

	var dps func(x int)
	dps = func(x int) {
		if x == len(c)-1 {
			result = append(result, string(c))
			return
		}
		sets := make(map[byte]struct{}, 0)
		for i := x; i < len(c); i++ {
			if _, ok := sets[c[i]]; ok {
				continue
			}
			sets[c[i]] = struct{}{}
			c[i], c[x] = c[x], c[i]
			dps(x + 1)
			c[i], c[x] = c[x], c[i]
		}
	}
	dps(0)
	return result
}