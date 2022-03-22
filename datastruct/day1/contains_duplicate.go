package day1

import "sort"

func ContainsDuplicate1(nums []int) bool {
	sets := make(map[int]struct{}, 0)

	for _, n := range nums {
		if ok := lookupSets(sets, n); ok {
			return true
		}
	}

	return false
}

func ContainsDuplicate2(nums []int) bool {
	sort.Ints(nums)

	length := len(nums)
	if length < 2 {
		return false
	}

	i := 0
	j := 1
	for j < length {
		if nums[i] == nums[j] {
			return true
		}

		i += 1
		j += 1
	}

	return false
}

func lookupSets(sets map[int]struct{}, k int) bool {
	_, ok := sets[k]
	if ok {
		return true
	}
	sets[k] = struct{}{}
	return false
}
