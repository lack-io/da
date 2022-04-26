package day22

func singleNumber(nums []int) int {
	x, y := 0, 0
	for _, i := range nums {
		x = x ^ i & ^y
		y = y ^ i & ^x
	}
	return x
}
