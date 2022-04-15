package queue

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q Queue) String() string {
	out := make([]string, 0)

	p := q.head
	for p != nil {
		out = append(out, strconv.Itoa(p.Val))
		p = p.Next
	}

	return strings.Join(out, "->")
}

func (q *Queue) Push(val int) {
	node := &Node{Val: val}
	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.Next = node
	q.tail = node
}

func (q *Queue) Pop() (int, error) {
	if q.head == nil {
		return 0, fmt.Errorf("empty queue")
	}
	val := q.head.Val
	q.head = q.head.Next
	return val, nil
}
