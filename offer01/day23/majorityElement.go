package day23

import "sort"

func majorityElement(nums []int) int {
	sets := make(map[int]int, 0)

	for _, n := range sets {
		sets[n] += 1
	}

	n := 0
	max := 0
	for k, v := range sets {
		if k > max {
			max = k
			n = v
		}
	}
	return n
}

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement2(nums []int) int {
	x := nums[0]
	sum := 0

	for _, n := range nums {
		if sum == 0 {
			x = n
		} else {
			if n == x {
				sum += 1
			} else {
				sum -= 1
			}
		}
	}

	return x
}
