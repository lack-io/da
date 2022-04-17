package day13

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1

	out := make([]int, 2)
	for i <= j {
		if nums[i]+nums[j] == target {
			out[0], out[1] = nums[i], nums[j]
			break
		}
		if nums[i]+nums[j] > target {
			j -= 1
		} else {
			i += 1
		}
	}

	return out
}
