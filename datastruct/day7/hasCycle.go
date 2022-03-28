package day7

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	out := make([]int, 0)

	head := l
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return fmt.Sprintf("%v", out)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func HasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == head {
		return true
	}

	m := make(map[*ListNode]struct{})
	for ; head.Next != nil; head = head.Next {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
	}

	return false
}
