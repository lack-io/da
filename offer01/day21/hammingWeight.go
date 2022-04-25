package day21

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num & uint32(1))
		num >>= 1
	}
	return count
}
