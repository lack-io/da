package sort

func QuickSort(nums []int) []int {
	quickSort(&nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums *[]int, start, end int) {
	if start < end {
		key := (*nums)[start]
		i := start
		for j := start + 1; j <= end; j++ {
			if key > (*nums)[j] {
				i += 1
				swap(nums, j, i)
			}
		}
		(*nums)[start] = (*nums)[i]
		(*nums)[i] = key
		quickSort(nums, start, i-1)
		quickSort(nums, i+1, end)
	}
}
