package day8

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

func ReverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var slow *ListNode
	fast := head
	for fast != nil {
		next := fast.Next
		fast.Next = slow
		slow = fast
		fast = next
	}
	return slow
}

func ReverseList2(head *ListNode) *ListNode {
	return reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
