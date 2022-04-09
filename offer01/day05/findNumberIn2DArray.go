package day05

func findNumberIn2DArray1(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, n := range row {
			if n == target {
				return true
			}
		}
	}
	return false
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if len(row) > 0 && row[0] > target {
			return false
		}
		for _, n := range row {
			if n > target {
				break
			}
			if n == target {
				return true
			}
		}
	}
	return false
}
