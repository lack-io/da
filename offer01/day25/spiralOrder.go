package day25

// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
func spiralOrder(matrix [][]int) []int {
	out := make([]int, 0)
	row, col := len(matrix), len(matrix[0])
	traversal(matrix, 0, col, 0, row, &out)
	return out
}

func traversal(matrix [][]int, rowi, rowj, coli, cowj int, out *[]int) {
	if rowj-rowi < 1 || cowj-coli < 1 {
		return
	}
	for i := rowi; i < rowj; i++ {
		*out = append(*out, matrix[coli][i])
	}
	for i := coli + 1; i < cowj; i++ {
		*out = append(*out, matrix[i][rowj-1])
	}
	if cowj-1 != coli {
		for i := rowj - 2; i >= rowi; i-- {
			*out = append(*out, matrix[cowj-1][i])
		}
	}
	if rowj-1 != rowi {
		for i := cowj - 2; i >= coli+1; i-- {
			*out = append(*out, matrix[i][rowi])
		}
	}
	traversal(matrix, rowi+1, rowj-1, coli+1, cowj-1, out)
}
