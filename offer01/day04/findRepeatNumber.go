package day04

func findRepeatNumber(nums []int) int {
	set := make(map[int]struct{}, 0)

	num := 0
	for _, n := range nums {
		if _, ok := set[n]; ok {
			num = n
			break
		}
		set[n] = struct{}{}
	}

	return num
}

func findRepeatNumber1(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i += 1
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}

	return -1
}
