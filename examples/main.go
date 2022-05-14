package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	fmt.Println(a & 1 << 2)
}

func Decimal(value float64, n int) float64 {
	p := 1e1
	for i := 0; i < n-1; i++ {
		p *= 1e1
	}

	return math.Trunc(value*p+0.5) * 1 / p
}
