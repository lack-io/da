package day29

// https://leetcode-cn.com/problems/chou-shu-lcof/
// i=1,j=2,k=3
// -> i=2,j=2,k=3 , r = i * j = 4
// -> i=2,j=2,k=5,  r = 5
func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, 0)
	dp = append(dp, 0)

	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(n2, min(n3, n5))
		if dp[i] == n2 {
			a += 1
		}
		if dp[i] == n3 {
			b += 1
		}
		if dp[i] == n5 {
			c += 1
		}
	}
	return dp[n-1]
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}
