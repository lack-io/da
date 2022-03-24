package day3

import "sort"

// 哈希法
func Intersect1(nums1 []int, nums2 []int) []int {
	out := make([]int, 0)
	sets := make(map[int]int, 0)
	for _, n := range nums1 {
		if _, ok := sets[n]; ok {
			sets[n] += 1
		} else {
			sets[n] = 1
		}
	}

	for _, n := range nums2 {
		if _, ok := sets[n]; ok {
			out = append(out, n)
			sets[n] -= 1
			if sets[n] == 0 {
				delete(sets, n)
			}
		}
	}

	return out
}

// 双指针法
func Intersect2(nums1 []int, nums2 []int) []int {
	outs := make([]int, 0)

	sort.Ints(nums1)
	sort.Ints(nums2)

	nums1Len := len(nums1)
	nums2Len := len(nums2)
	p1, p2 := 0, 0

	for {
		if p1 == nums1Len {
			break
		}
		if p2 >= nums2Len {
			break
		}

		m, n := nums1[p1], nums2[p2]

		if m == n {
			outs = append(outs, m)
			p1 += 1
			p2 += 1
		} else if m > n {
			p2 += 1
		} else {
			p1 += 1
		}
	}

	return outs
}
