package day2

import (
	"reflect"
	"testing"
)

func TestMerge1(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "TestMerge1-1", args: struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3}, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "TestMerge1-2", args: struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}{nums1: []int{1}, m: 1, nums2: []int{}, n: 0}, want: []int{1}},
		{name: "TestMerge1-3", args: struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1}, want: []int{1}},
		{name: "TestMerge1-4", args: struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}{nums1: []int{2, 0}, m: 1, nums2: []int{1}, n: 1}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge1(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("Merge1() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}

func TestMerge2(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "TestMerge2-1", args: struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3}, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "TestMerge2-2", args: struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}{nums1: []int{1}, m: 1, nums2: []int{}, n: 0}, want: []int{1}},
		{name: "TestMerge2-3", args: struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1}, want: []int{1}},
		{name: "TestMerge2-4", args: struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}{nums1: []int{2, 0}, m: 1, nums2: []int{1}, n: 1}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge2(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("Mrege2() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
