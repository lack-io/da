package day

import "math"

func largestTriangleArea(points [][]int) float64 {
	maxArea := float64(0)
	for i, p := range points {
		for j, q := range points[:i] {
			for _, r := range points[:j] {
				maxArea = math.Max(maxArea, area(p[0], p[1], q[0], q[1], r[0], r[1]))
			}
		}
	}
	return maxArea
}

func area(x1, y1, x2, y2, x3, y3 int) float64 {
	return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) * 1 / 2
}
