package day

func threeSum(nums []int) [][]int {
	out := make([][]int, 0)

	length := len(nums)
	if length < 3 {
		return out
	}

	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					out = append(out, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return out
}
