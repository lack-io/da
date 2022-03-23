package day2

func TwoSum1(nums []int, target int) []int {
	length := len(nums)

	if length < 2 {
		return nil
	}

	out := make([]int, 0)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				out = []int{i, j}
			}
		}
	}

	return out
}

func TwoSum2(nums []int, target int) []int {
	ht := map[int]int{}
	for i, n := range nums {
		if p, ok := ht[target-n]; ok {
			return []int{p, i}
		}
		ht[n] = i
	}

	return nil
}
