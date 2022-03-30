package day9

import (
	"fmt"
)

func IsValid1(s string) bool {
	stack := NewStack()
	brackets := map[int32]int32{')': '(', '}': '{', ']': '['}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.Push(c)
		}
		if c == ')' || c == '}' || c == ']' {
			r, err := stack.Pop()
			if err != nil {
				return false
			}
			if brackets[c] != r {
				return false
			}
		}
	}
	return stack.Len() == 0
}

type Stack struct {
	data []int32
}

func NewStack() *Stack {
	return &Stack{data: []int32{}}
}

func (s *Stack) Push(c int32) {
	s.data = append(s.data, c)
}

func (s *Stack) Pop() (int32, error) {
	if s.Len() == 0 {
		return -1, fmt.Errorf("empty stack")
	}
	c := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return c, nil
}

func (s *Stack) Len() int {
	return len(s.data)
}
