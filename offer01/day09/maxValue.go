package day09

func maxValue(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else if j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + max(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[m-1][n-1]
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
