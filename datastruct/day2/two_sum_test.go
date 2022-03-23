package day2

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "TestTwoSum1-1", args: struct {
			nums   []int
			target int
		}{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "TestTwoSum1-2", args: struct {
			nums   []int
			target int
		}{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{name: "TestTwoSum1-3", args: struct {
			nums   []int
			target int
		}{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSum2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "TestTwoSum2-1", args: struct {
			nums   []int
			target int
		}{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "TestTwoSum2-2", args: struct {
			nums   []int
			target int
		}{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{name: "TestTwoSum2-3", args: struct {
			nums   []int
			target int
		}{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
