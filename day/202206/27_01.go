package _02206

func isSubseq(s, t string) bool {
	ptS := 0
	for ptT := range t {
		if s[ptS] == t[ptT] {
			if ptS++; ptS == len(s) {
				return true
			}
		}
	}
	return false
}

func findLUSlength(strs []string) int {
	ans := -1
next:
	for i, s := range strs {
		for j, t := range strs {
			if i != j && isSubseq(s, t) {
				continue next
			}
		}
		if len(s) > ans {
			ans = len(s)
		}
	}
	return ans
}
