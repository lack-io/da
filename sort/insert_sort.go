package sort

func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] > nums[j-1] {
				break
			}
			swap(&nums, j-1, j)
		}
	}
	return nums
}
