package day

func findClosest(words []string, word1 string, word2 string) int {
	dict := make(map[string][]int, 0)

	for i, s := range words {
		_, ok := dict[s]
		if !ok {
			dict[s] = make([]int, 0)
		}
		dict[s] = append(dict[s], i)
	}

	p, q := dict[word1], dict[word2]

	ans := len(words)
	for i := 0; i < len(q); i++ {
		for j := 0; j < len(p); j++ {
			ans = min(ans, abs(q[i]-p[j]))
		}
	}

	return ans
}

func findClosest1(words []string, word1 string, word2 string) int {
	p, q := 0, 0

	ans := len(words)
	for i, s := range words {
		if s == word1 {
			p = i
		}
		if s == word2 {
			q = i
		}

		ans = min(ans, abs(p-q))
	}

	return ans
}
