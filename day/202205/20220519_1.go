package _02205

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {

	length := len(nums)
	dp := make([]int, length+1)
	dp[0] = math.MaxInt

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j == i {
				continue
			}
			dp[i+1] += abs(nums[j] - nums[i])
		}

		dp[i+1] = min(dp[i+1], dp[i])
	}

	return dp[length]
}

func minMoves2_1(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	ans := 0
	for left < right {
		ans += nums[right] - nums[left]
		left += 1
		right -= 1
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
