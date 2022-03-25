package day4

func Generate1(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[i][j] = 1
			} else {
				ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
	}

	return ans
}
