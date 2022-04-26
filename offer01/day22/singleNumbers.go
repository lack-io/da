package day22

func singleNumbers(nums []int) []int {
	x, y, m, n := 0, 0, 0, 0

	for _, i := range nums {
		n ^= i
	}

	for n&m == 0 {
		m <<= 1
	}

	for _, i := range nums {
		if i&m != 0 {
			x ^= i
		} else {
			y ^= i
		}
	}

	return []int{x, y}
}
