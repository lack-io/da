package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()

	q.Push(1)
	q.Push(2)
	out := q.String()
	if out != "1->2" {
		t.Errorf("push failed: got %v, want %v", out, "1->2")
	}

	v, err := q.Pop()
	if err != nil {
		t.Errorf("pop failed: %v", err)
	}
	if v != 1 {
		t.Errorf("pop failed: got %d, want %v", v, 1)
	}

	v, err = q.Pop()
	if err != nil {
		t.Errorf("pop failed: %v", err)
	}
	if v != 2 {
		t.Errorf("pop failed: got %d, want %v", v, 2)
	}

	v, err = q.Pop()
	if err == nil {
		t.Errorf("pop failed: queue is not empty")
	}
}
