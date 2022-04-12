package day08

func numWays(n int) int {
	if n < 2 {
		return 1
	}

	p, q := 1, 1
	var result int
	for i := 2; i <= n; i++ {
		result = p + q
		result %= 1e9 + 7

		p, q = q, result
	}

	return result
}
