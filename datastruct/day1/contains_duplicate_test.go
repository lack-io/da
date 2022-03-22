package day1

import "testing"

func TestContainsDuplicate1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestContainsDuplicate1-1", args: struct{ nums []int }{nums: []int{1, 2, 3, 1}}, want: true},
		{name: "TestContainsDuplicate1-2", args: struct{ nums []int }{nums: []int{1, 2, 3, 4}}, want: false},
		{name: "TestContainsDuplicate1-3", args: struct{ nums []int }{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate1(tt.args.nums); got != tt.want {
				t.Errorf("ContainsDuplicate1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsDuplicate2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestContainsDuplicate2-1", args: struct{ nums []int }{nums: []int{1, 2, 3, 1}}, want: true},
		{name: "TestContainsDuplicate2-2", args: struct{ nums []int }{nums: []int{1, 2, 3, 4}}, want: false},
		{name: "TestContainsDuplicate2-3", args: struct{ nums []int }{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate2(tt.args.nums); got != tt.want {
				t.Errorf("ContainsDuplicate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
