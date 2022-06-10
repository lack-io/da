package _02206

import (
	"math/rand"
	"sort"
)

type __Solution struct {
	rects [][]int
	sum   []int
}

func __Constructor(rects [][]int) __Solution {
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
	}
	return __Solution{rects: rects, sum: sum}
}

func (s *__Solution) Pick() []int {
	k := rand.Intn(s.sum[len(s.sum)-1])
	idx := sort.SearchInts(s.sum, k+1) - 1
	r := s.rects[idx]
	a, b, y := r[0], r[1], r[3]
	da := (k - s.sum[idx]) / (y - b + 1)
	db := (k - s.sum[idx]) % (y - b + 1)
	return []int{a + da, b + db}
}
