package day30

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(nums []int) int {
	tmp := make([]int, len(nums))
	return mergeSort(0, len(nums)-1, nums, tmp)
}

func mergeSort(l, r int, nums, tmp []int) int {
	if l >= r {
		return 0
	}
	m := (l + r) / 2
	res := mergeSort(l, m, nums, tmp) + mergeSort(m+1, r, nums, tmp)

	i, j := l, m+1
	for k := l; k <= r; k++ {
		tmp[k] = nums[k]
	}
	for k := l; k <= r; k++ {
		if i == m+1 {
			nums[k] = tmp[j]
			j += 1
		} else if j == r+1 || tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i += 1
		} else {
			nums[k] = tmp[j]
			res += m - i + 1 // 统计逆序对
			j += 1
		}
	}
	return res
}
