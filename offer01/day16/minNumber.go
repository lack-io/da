package day16

import (
	"sort"
	"strconv"
)

type Ints []int

func (in Ints) Len() int {
	return len(in)
}

func (in *Ints) Less(i, j int) bool {
	a, b := (*in)[i], (*in)[j]

	if a == 0 {
		return true
	}
	if b == 0 {
		return false
	}

	aSum := a
	for n := 1; n < intBit(b); n++ {
		aSum *= 10
	}
	aSum += b

	bSum := b
	for n := 1; n < intBit(a); n++ {
		bSum *= 10
	}
	bSum += a

	return aSum < bSum
}

func (in *Ints) Swap(i, j int) {
	(*in)[i], (*in)[j] = (*in)[j], (*in)[i]
}

func intBit(num int) int {
	n := 1
	for num > 0 {
		n += 1
		num /= 10
	}
	return n
}

func minNumber(nums []int) string {
	numsInts := Ints(nums)
	sort.Sort(&numsInts)

	s := ""
	for _, i := range numsInts {
		s += strconv.Itoa(i)
	}
	return s
}
