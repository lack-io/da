package day

func lengthOfLongestSubstring(s string) int {

	n := len(s)
	if n == 0 {
		return 0
	}

	max := 1

	start := 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= start; j-- {
			if s[j] == s[i] {
				start = j + 1
			}
		}
		t := i - start + 1
		if t > max {
			max = t
		}
	}

	return max
}
