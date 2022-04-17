package day12

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2

	out := &ListNode{}
	head := out

	for p1 != nil || p2 != nil {

		if p1 == nil {
			out.Next = p2
			break
		}

		if p2 == nil {
			out.Next = p1
			break
		}

		if p1.Val > p2.Val {
			out.Next = &ListNode{Val: p2.Val}
			out = out.Next
			p2 = p2.Next
		} else {
			out.Next = &ListNode{Val: p1.Val}
			out = out.Next
			p1 = p1.Next
		}

	}

	return head.Next
}
