package day3

// 暴力法，超出时间限制
func MaxProfit1(prices []int) int {
	maxProfit := 0

	for i, n := range prices {
		for j := i + 1; j < len(prices); j++ {
			if maxProfit < prices[j]-n {
				maxProfit = prices[j] - n
			}
		}
	}

	return maxProfit
}

// 动态规划法
func MaxProfit2(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
