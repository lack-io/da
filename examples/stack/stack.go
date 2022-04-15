package stack

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s Stack) String() string {
	out := make([]string, 0)

	p := s.top
	for p != nil {
		out = append(out, strconv.Itoa(p.Val))
		p = p.Next
	}

	return strings.Join(out, "->")
}

func (s *Stack) Push(val int) {
	node := &Node{Val: val}
	if s.top == nil {
		s.top = node
		return
	}
	node.Next = s.top
	s.top = node
}

func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("empty stack")
	}
	val := s.top.Val
	s.top = s.top.Next
	return val, nil
}
