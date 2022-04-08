package day04

func missingNumber(nums []int) int {
	p1, p2 := 0, len(nums)
	if nums[p1] != p1 {
		return 0
	}
	if nums[len(nums)-1] != p2 {
		return p2
	}

	n := 0
	for p1 < p2 {
		if nums[p1] != p1 {
			n = p1
			break
		}

		if nums[p2] != p2 {
			n = p2
			break
		}

		p1 += 1
		p2 -= 1
	}

	return n
}

func missingNumber1(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}
