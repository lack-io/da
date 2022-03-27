package day6

func CanConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	mSet := make(map[int32]int)
	for _, c := range magazine {
		mSet[c] += 1
	}

	for _, k := range ransomNote {
		if v, ok := mSet[k]; !ok {
			return false
		} else {
			if v == 0 {
				return false
			}
			mSet[k] -= 1
		}
	}
	return true
}
