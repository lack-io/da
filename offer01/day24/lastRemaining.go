package day24

func lastRemaining(n, m int) int {
	x := 0
	for i := 2; i <= n; i++ {
		x = (x + m) % i
	}
	return x
}
