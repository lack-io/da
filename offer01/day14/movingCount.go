package day14

func movingCount(m int, n int, k int) int {
	path := make([][]int, m)
	count := 0

	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}
	moveDFS(m, n, k, 0, 0, &path, &count)
	return count
}

func moveDFS(m, n, k, i, j int, path *[][]int, count *int) {
	if i >= m || i < 0 || j >= n || j < 0 || (*path)[i][j] == 1 {
		return
	}

	iSum, jSum := nadd(i), nadd(j)
	if iSum+jSum > k {
		return
	}

	(*path)[i][j] = 1
	if iSum+jSum <= k {
		*count += 1
	}
	moveDFS(m, n, k, i+1, j, path, count)
	moveDFS(m, n, k, i, j+1, path, count)
}

func nadd(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
