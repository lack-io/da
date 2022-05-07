package day29

import (
	"math"
)

// https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/
// n = 1, 1 1 1  1  1  1
// n = 2, 1 2 3  4  5  6  5  4  3  2  1
// n = 3, 1 3 6 10 15 21 25 27 27 25 21 15 10 6 3 1
func dicesProbability(n int) []float64 {
	total := math.Pow(6, float64(n))
	dp := make([][]int, n)
	for i := 1; i <= n; i++ {
		dp[i-1] = make([]int, i*6-(i-1))
	}
	dp[0] = []int{1, 1, 1, 1, 1, 1}

	for i := 1; i < n; i++ {
		dp[i][0] = 1
		dp[i][len(dp[i])-1] = 1
		for j := 1; j < len(dp[i])-1; j++ {
			dp[i][j] = dp[i][j-1]
			if j >= 6 {
				dp[i][j] -= dp[i-1][j-6]
			}
			if j < len(dp[i-1]) {
				dp[i][j] += dp[i-1][j]
			}
		}
	}

	out := make([]float64, len(dp[n-1]))
	for i := 0; i < len(dp[n-1]); i++ {
		out[i] = decimal(float64(dp[n-1][i])/total, 5)
	}
	return out
}

// decimal 保留小数点个数
func decimal(value float64, n int) float64 {
	p := 1e1
	for i := 0; i < n-1; i++ {
		p *= 1e1
	}

	return math.Trunc(value*p+0.5) * 1 / p
}
