package day30

import (
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
func printNumbers(n int) []int {
	res := make([]int, int(math.Pow10(n)-1))
	count := 0
	var dfs func(index int, num []byte, digit int)
	dfs = func(index int, num []byte, digit int) {
		if index == digit {
			tmp, _ := strconv.Atoi(string(num))
			res[count] = tmp
			count++
			return
		}
		for i := '0'; i <= '9'; i++ {
			num[index] = byte(i)
			dfs(index+1, num, digit)
		}
	}
	for digit := 1; digit <= n; digit++ {
		for first := '1'; first <= '9'; first++ {
			num := make([]byte, digit)
			num[0] = byte(first)
			dfs(1, num, digit)
		}
	}
	return res
}
