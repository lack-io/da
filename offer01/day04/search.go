package day04

func search(nums []int, target int) int {
	p1, p2 := 0, len(nums)-1
	for p1 <= p2 {
		if nums[p1] == target && nums[p2] == target {
			return p2 - p1 + 1
		}
		if nums[p1] != target {
			p1 += 1
		}
		if nums[p2] != target {
			p2 -= 1
		}
	}

	return 0
}
