package day31

import "testing"

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_findNthDigit", struct{ n int }{n: 3}, 3},
		{"Test_findNthDigit-1", struct{ n int }{n: 11}, 0},
		{"Test_findNthDigit-2", struct{ n int }{n: 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
