package day21

func add(a, b int) int {
	// 当进位为 0 时跳出
	for b != 0 {
		// c = 进位
		c := (a & b) << 1
		// a = 非进位和
		a ^= b
		// b = 进位
		b = c
	}
	return a
}
