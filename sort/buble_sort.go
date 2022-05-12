package sort

func swap(nums *[]int, i, j int) {
	if i != j {
		(*nums)[i] ^= (*nums)[j]
		(*nums)[j] ^= (*nums)[i]
		(*nums)[i] ^= (*nums)[j]
	}
}

func BubleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[j-1] {
				swap(&nums, j, j-1)
			}
		}
	}
	return nums
}
