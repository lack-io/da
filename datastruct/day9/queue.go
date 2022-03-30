package day9

type MyQueue struct {
	head *stack
	tail *stack
}

func Constructor() MyQueue {
	return MyQueue{head: newStack(), tail: newStack()}
}

func (this *MyQueue) Push(x int) {
	this.tail.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	if !this.head.IsEmpty() {
		return this.head.Pop()
	}
	for !this.tail.IsEmpty() {
		this.head.Push(this.tail.Pop())
	}
	return this.head.Pop()
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	if !this.head.IsEmpty() {
		return this.head.Peek()
	}
	for !this.tail.IsEmpty() {
		this.head.Push(this.tail.Pop())
	}
	return this.head.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.head.IsEmpty() && this.tail.IsEmpty()
}

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{data: []int{}}
}

func (s *stack) Push(c int) {
	s.data = append(s.data, c)
}

func (s *stack) Pop() int {
	if s.Size() == 0 {
		return -1
	}
	c := s.data[s.Size()-1]
	s.data = s.data[:s.Size()-1]
	return c
}

func (s *stack) Peek() int {
	if s.Size() == 0 {
		return -1
	}
	return s.data[s.Size()-1]
}

func (s *stack) Size() int {
	return len(s.data)
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}
