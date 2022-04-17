package day13

func exchange(nums []int) []int {
	result := make([]int, len(nums))

	m, n := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[n] = nums[i]
			n -= 1
		} else {
			nums[m] = nums[i]
			m += 1
		}
	}

	return result
}

func exchange1(nums []int) []int {
	m, n := 0, len(nums)-1

	for m <= n {
		if nums[m]%2 == 0 && nums[n]%2 != 0 {
			nums[m], nums[n] = nums[n], nums[m]
			m += 1
			n -= 1
			continue
		}

		if nums[m]%2 != 0 {
			m += 1
		}
		if nums[n]%2 == 0 {
			n -= 1
		}
	}
	return nums
}
