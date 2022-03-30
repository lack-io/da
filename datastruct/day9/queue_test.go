package day9

import (
	"testing"
)

func TestMyQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	n := q.Peek()
	if n != 1 {
		t.Fatalf("Peek: got = %d, want = %d", n, 1)
	}
	n = q.Pop()
	if n != 1 {
		t.Fatalf("Pop: got = %d, want = %d", n, 1)
	}
	if q.Empty() {
		t.Fatalf("Empty: got = %v, want = %v", q.Empty(), false)
	}
}
