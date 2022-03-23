package day1

import "testing"

func TestMaxSubArray1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestMaxSubArray1-1", args: struct{ nums []int }{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, want: 6},
		{name: "TestMaxSubArray1-2", args: struct{ nums []int }{nums: []int{1}}, want: 1},
		{name: "TestMaxSubArray1-3", args: struct{ nums []int }{nums: []int{5, 4, -1, 7, 8}}, want: 23},
		{name: "TestMaxSubArray1-3", args: struct{ nums []int }{nums: []int{-1}}, want: -1},
		{name: "TestMaxSubArray1-3", args: struct{ nums []int }{nums: []int{1, 2, -1, -2, 2, 1, -2, 1}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray1(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArray1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubArray2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestMaxSubArray2-1", args: struct{ nums []int }{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, want: 6},
		{name: "TestMaxSubArray2-2", args: struct{ nums []int }{nums: []int{1}}, want: 1},
		{name: "TestMaxSubArray2-3", args: struct{ nums []int }{nums: []int{5, 4, -1, 7, 8}}, want: 23},
		{name: "TestMaxSubArray2-3", args: struct{ nums []int }{nums: []int{-1}}, want: -1},
		{name: "TestMaxSubArray2-3", args: struct{ nums []int }{nums: []int{1, 2, -1, -2, 2, 1, -2, 1}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}
