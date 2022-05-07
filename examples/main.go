package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Decimal(float64(1)/math.Pow(6, 2), 5))
}

func Decimal(value float64, n int) float64 {
	p := 1e1
	for i := 0; i < n-1; i++ {
		p *= 1e1
	}

	return math.Trunc(value*p+0.5) * 1 / p
}
