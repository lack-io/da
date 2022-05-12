package day

import (
	"testing"
)

func Test_Codec(t *testing.T) {
	codec := Codec{}
	root := codec.deserialize("[2,1,3]")
	t.Log(root)

	out := codec.deserialize(codec.serialize(root))
	t.Log(out)
}

func Test_AddTwoNumbers(t *testing.T) {
	l := addTwoNumbers(&ListNode{Val: 7}, &ListNode{Val: 7})
	t.Log(l)
}

func Test_minDeletionSize(t *testing.T) {
	strs := []string{"zyx", "wvu", "tsr"}
	count := minDeletionSize(strs)
	if count != 3 {
		t.Fatalf("got=%d, want=%d", count, 3)
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	m := findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6, 7, 8})
	t.Log(m)
}
