package day31

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

	p := 1e9 + 7
	res := 1
	for n > 4 {
		res = (res * 3) % int(p)
		n -= 3
	}
	res *= n % int(p)
	return res % int(p)
}
