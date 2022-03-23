package day2

import "sort"

func Merge1(nums1 []int, m int, nums2 []int, n int) {
	// 可以简化成
	//	copy(nums1[m:], nums2)
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)
}

// 双指针
func Merge2(nums1 []int, m int, nums2 []int, n int) {

	out := make([]int, 0, m+n)

	p1, p2 := 0, 0
	for {
		if p1 == m {
			out = append(out, nums2[p2:]...)
			break
		}
		if p2 == n {
			out = append(out, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			out = append(out, nums1[p1])
			p1++
		} else {
			out = append(out, nums2[p2])
			p2++
		}
	}

	copy(nums1, out)
}
