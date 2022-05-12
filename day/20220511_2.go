package day

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2

	var root *ListNode = nil
	var t *ListNode = nil

	set := func(m, n *ListNode, bit *int) {
		val := *bit
		if m != nil {
			val += m.Val
		}
		if n != nil {
			val += n.Val
		}
		*bit = 0
		if val >= 10 {
			*bit = 1
		}
		if root == nil {
			root = &ListNode{Val: val % 10}
			t = root
		} else {
			t.Next = &ListNode{Val: val % 10}
			t = t.Next
		}
	}

	bit := 0
	for {
		if p == nil {
			for ; q != nil; q = q.Next {
				set(p, q, &bit)
			}
			break
		}

		if q == nil {
			for ; p != nil; p = p.Next {
				set(p, q, &bit)
			}
			break
		}

		set(p, q, &bit)
		p = p.Next
		q = q.Next
	}
	if bit > 0 {
		t.Next = &ListNode{Val: bit}
		t = t.Next
	}

	return root
}
