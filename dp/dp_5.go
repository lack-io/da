package dp

// https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	// 1. 定义数组元素的含义
	// 2. 找出元素之间的关系式
	// 3. 找出初始值
	var pralindrome func(text string) string
	pralindrome = func(text string) string {
		if len(text) == 1 {
			return text
		}
		length := len(text)
		p, q := 0, length-1
		for p <= q {
			if text[p] == text[q] {
				p += 1
				q -= 1
			} else {
				return pralindrome(text[1:])
			}
		}
		return text
	}

	n := len(s)
	dp := make([]string, n)
	dp[0] = string(s[0])
	for i := 1; i < n; i++ {
		dp[i] = pralindrome(s[:i+1])
		if len(dp[i-1]) > len(dp[i]) {
			dp[i] = dp[i-1]
		}
	}

	return dp[n-1]
}
