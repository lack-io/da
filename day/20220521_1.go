package day

func repeatedNTimes(nums []int) int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
		if m[nums[i]] >= 2 {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
