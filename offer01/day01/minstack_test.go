package day01

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.Pop()

	top := s.Top()
	if top != 0 {
		t.Fatalf("top failed: got = %d, want = %d", top, 0)
	}

	min := s.Min()
	if min != -2 {
		t.Fatalf("min failed: got = %d, want = %d", min, -2)
	}
}

func TestMinStack1(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-1)
	s.Pop()

	min := s.Min()
	if min != -2 {
		t.Fatalf("min failed: got = %d, want = %d", min, -2)
	}

	top := s.Top()
	if top != 0 {
		t.Fatalf("top failed: got = %d, want = %d", top, 0)
	}
}
