package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	out := s.String()
	if out != "2->1" {
		t.Errorf("push failed: got %v, want %v", out, "2->1")
	}

	v, err := s.Pop()
	if err != nil {
		t.Errorf("pop failed: %v", err)
	}
	if v != 2 {
		t.Errorf("pop failed: got %d, want %v", v, 2)
	}

	v, err = s.Pop()
	if err != nil {
		t.Errorf("pop failed: %v", err)
	}
	if v != 1 {
		t.Errorf("pop failed: got %d, want %v", v, 1)
	}

	v, err = s.Pop()
	if err == nil {
		t.Errorf("pop failed: stack is not empty")
	}
}
