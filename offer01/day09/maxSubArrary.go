package day09

func maxSubArray(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < length; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		result = max(result, dp[i])
	}

	return result
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func maxSubArray1(nums []int) int {
	length := len(nums)
	result := nums[0]

	for i := 1; i < length; i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}

		result = max(result, nums[i])
	}

	return result
}
