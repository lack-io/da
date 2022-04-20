package day16

func isStraight(nums []int) bool {
	sets := make(map[int]struct{})

	min := 13
	max := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if _, ok := sets[nums[i]]; ok {
			return false
		}
		sets[nums[i]] = struct{}{}

		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	return max-min < 5
}
