package day

import (
	"fmt"
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

func Test_isUnique(t *testing.T) {
	fmt.Println(isUnique("abc"))
}

func Test_Validate(t *testing.T) {
	t.Log(Validate("5"))
	t.Log(Validate("Ooook"))
	t.Log(Validate("Hhhhh55"))
}

func Test_convert(t *testing.T) {
	//t.Log(convert("PAYPALISHIRING", 3))
	t.Log(convert("PLRAAIIIYASINPHG", 4))
}

func Test_reverse(t *testing.T) {
	t.Log(reverse(123))
}
