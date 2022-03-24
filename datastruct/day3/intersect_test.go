package day3

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersect1(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "TestIntersect1-1", args: struct {
			num1 []int
			num2 []int
		}{num1: []int{1, 2, 2, 1}, num2: []int{2, 2}}, want: []int{2, 2}},
		{name: "TestIntersect1-2", args: struct {
			num1 []int
			num2 []int
		}{num1: []int{4, 9, 5}, num2: []int{9, 4, 9, 8, 4}}, want: []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersect1(tt.args.num1, tt.args.num2)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect2(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "TestIntersect2-1", args: struct {
			nums1 []int
			nums2 []int
		}{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2, 2}},
		{name: "TestIntersect2-2", args: struct {
			nums1 []int
			nums2 []int
		}{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersect2(tt.args.nums1, tt.args.nums2)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect2() = %v, want %v", got, tt.want)
			}
		})
	}
}
