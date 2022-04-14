package day10

func translateNum(num int) int {
	nums := make([]int, 0)

	if num == 0 {
		return 1
	}

	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	length := len(nums)
	dp := make([]int, length)
	dp[length-1] = 1

	for i := length - 2; i >= 0; i-- {
		if translatable(nums[i+1], nums[i]) {
			if length-i < 3 {
				dp[i] = dp[i+1] + 1
			} else {
				dp[i] = dp[i+1] + dp[i+2]
			}
		} else {
			dp[i] = dp[i+1]
		}
	}

	return dp[0]
}

func translatable(m, n int) bool {
	n = m*10 + n
	return n >= 10 && n <= 25
}
