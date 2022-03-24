package day3

import "testing"

func TestMaxProfit1(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestMaxProfit1-1", args: struct{ prices []int }{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
		{name: "TestMaxProfit1-1", args: struct{ prices []int }{prices: []int{7, 6, 4, 3, 1}}, want: 0},
		{name: "TestMaxProfit1-1", args: struct{ prices []int }{prices: []int{2, 4, 1}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit1(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestMaxProfit2-1", args: struct{ prices []int }{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
		{name: "TestMaxProfit2-1", args: struct{ prices []int }{prices: []int{7, 6, 4, 3, 1}}, want: 0},
		{name: "TestMaxProfit2-1", args: struct{ prices []int }{prices: []int{2, 4, 1}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
