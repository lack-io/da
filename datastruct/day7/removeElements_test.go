package day7

import (
	"reflect"
	"testing"
)

func TestRemoveElements1(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "TestRemoveElements1-1",
			args: struct {
				head *ListNode
				val  int
			}{head: &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}, val: 6},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElements1(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElements1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElements2(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "TestRemoveElements2-1",
			args: struct {
				head *ListNode
				val  int
			}{head: &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}, val: 6},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElements2(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElements2() = %v, want %v", got, tt.want)
			}
		})
	}
}
