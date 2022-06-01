package _02205

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	sum := 0

	for i := 0; i < m; i++ {
		sum += int(s1[i])
	}
	for i := 0; i < n; i++ {
		sum += int(s2[i])
	}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return sum - 2*dp[m][n]
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
