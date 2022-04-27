package heap

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()

	h.Push(1)
	h.Push(3)
	h.Push(2)

	top, err := h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if top != 3 {
		t.Fatalf("top: want = %d, got = %d", 3, top)
	}

	v, err := h.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if v != 3 {
		t.Fatalf("pop: want = %d, got = %d", 3, v)
	}

	top, err = h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if top != 2 {
		t.Fatalf("top: want = %d, got = %d", 2, top)
	}
}
