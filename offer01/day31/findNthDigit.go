package day31

import (
	"strconv"
)

// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
// 0-9			1 位 ---> 0 - 9		 9  个          9 * (1-1) * 10 * 1
// 10-99    	2 位 ---> 10 -       90 * 2  个     9 * (2-1) * 10 * 2
// 100-999  	3 位 ---> 1000 -     900 * 3 个
// 10000-9999   4 为 ---> 10000 -	 9000 * 4 个    9 * (k-1) * 10 * k
func findNthDigit(n int) int {
	// 1. 先确定所在范围
	// 2. 再确定在哪个数上
	// 3. 最后确定位数

	// 1.
	digit, digitNum, count := 1, 1, 9
	for n > count {
		n -= count
		digit++
		digitNum *= 10
		count = 9 * digit * digitNum
	}
	num := digitNum + (n-1)/digit

	index := (n - 1) % digit

	numStr := strconv.Itoa(num)
	return int(numStr[index] - '0')
}
