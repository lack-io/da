package day7

import "testing"

func TestHasCycle1(t *testing.T) {
	type args struct {
		head *ListNode
	}

	ln1 := &ListNode{Val: 3}
	ln2 := &ListNode{Val: 2}
	ln1.Next = ln2
	ln3 := &ListNode{Val: 0}
	ln2.Next = ln3
	ln4 := &ListNode{Val: -4}
	ln3.Next = ln4
	ln4.Next = ln2
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestHasCycle1-1", args: struct{ head *ListNode }{head: ln1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle1(tt.args.head); got != tt.want {
				t.Errorf("HasCycle1() = %v, want %v", got, tt.want)
			}
		})
	}
}
