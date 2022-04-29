package day24

func findContinuousSequence(target int) [][]int {
	i := 1   // 滑动窗口的左边界
	j := 1   // 滑动窗口的右边界
	sum := 0 // 滑动窗口中数字的和
	res := make([][]int, 0)

	for i <= target/2 {
		if sum < target {
			// 右边界向右移动
			sum += j
			j++
		} else if sum > target {
			// 左边界向右移动
			sum -= i
			i++
		} else {
			// 记录结果
			arr := make([]int, 0)
			for k := i; k < j; k++ {
				arr = append(arr, k)
			}
			res = append(res, arr)
			// 左边界向右移动
			sum -= i
			i++
		}
	}

	return res
}
