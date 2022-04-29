package day24

func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}

	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}
	res *= n
	return res
}
