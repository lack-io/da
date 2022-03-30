package day8

import (
	"reflect"
	"testing"
)

func TestReverseList1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "TestReverseList1-1",
			args: struct{ head *ListNode }{head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}},
			want: &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseList2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "TestReverseList1-1",
			args: struct{ head *ListNode }{head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}},
			want: &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList2() = %v, want %v", got, tt.want)
			}
		})
	}
}
