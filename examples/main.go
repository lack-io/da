package main

import (
	"fmt"
	"math"

	"github.com/lack-io/da/sort"
)

func main() {
	a := []int{1, 5, 2, 7, 10, 2, 234, 12, 431, 31, 112, 53, 123}
	fmt.Println(sort.InsertSort(a))
}

func Decimal(value float64, n int) float64 {
	p := 1e1
	for i := 0; i < n-1; i++ {
		p *= 1e1
	}

	return math.Trunc(value*p+0.5) * 1 / p
}
