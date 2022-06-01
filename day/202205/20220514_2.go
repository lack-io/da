package _02205

func isUnique(astr string) bool {
	bit := 0
	for _, c := range astr {
		if bit&(1<<(c-'a')) != 0 {
			return false
		}
		bit |= (1 << (c - 'a'))
	}
	return true
}
