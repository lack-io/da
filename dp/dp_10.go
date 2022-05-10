package dp

// https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/
func isMatch(s, p string) bool {
	m, n := len(s)+1, len(p)+1
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true

	for j := 2; j < n; j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && s[i-1] == p[j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && p[j-2] == '.' {
					dp[i][j] = true
				}
			} else {
				if dp[i-1][j-1] && s[i-1] == p[j-1] {
					dp[i][j] = true
				} else if dp[i-1][j-1] && p[j-1] == '.' {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[m-1][n-1]
}
