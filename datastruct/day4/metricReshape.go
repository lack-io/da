package day4

func MatrixReshape1(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m*n != r*c {
		return mat
	}

	outs := make([][]int, 0)
	line := make([]int, 0)
	lineIndex := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			line = append(line, mat[i][j])
			lineIndex += 1
			if lineIndex == c {
				outs = append(outs, line)
				lineIndex = 0
				line = make([]int, 0)
			}
		}
	}

	return outs
}

func MatrixReshape2(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		ans[i/c][i%c] = mat[i/n][i%n]
	}

	return ans
}
