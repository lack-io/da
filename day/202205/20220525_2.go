package _02205

func searchInsert(nums []int, target int) int {
	var search func(nums []int, start, end, target int) int
	search = func(nums []int, start, end, target int) int {
		if nums[start] >= target {
			return start
		}
		if nums[end] < target {
			return end + 1
		}
		if nums[end] == target {
			return end
		}

		if nums[start] < target && nums[end] > target && end-start == 1 {
			return end
		}

		mid := (start + end) / 2
		if nums[mid] >= target {
			return search(nums, start, mid, target)
		}
		return search(nums, mid, end, target)
	}

	return search(nums, 0, len(nums)-1, target)
}

func searchInsert1(nums []int, target int) int {
	length := len(nums)

	l, r := 0, length-1

	for l <= r {
		mid := l + (l+r)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
