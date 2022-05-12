package day

func minDeletionSize(strs []string) int {
	count := 0
	if len(strs) == 0 || len(strs[0]) == 0 {
		return count
	}
	row := len(strs)
	col := len(strs[0])
	for i := 0; i < col; i++ {
		last := uint8(0)
		for j := 0; j < row; j++ {
			if strs[j][i] < last {
				count += 1
				break
			}
			last = strs[j][i]
		}
	}

	return count
}
