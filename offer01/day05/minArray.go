package day05

func minArray(nums []int) int {
	min := nums[len(nums)-1]

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < min {
			min = nums[i]
			continue
		}
		if nums[i] > min {
			break
		}
	}

	return min
}
