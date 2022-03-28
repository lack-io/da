package day7

func RemoveElements1(head *ListNode, val int) *ListNode {

	p1 := head
	out := &ListNode{}
	p2 := out
	for p1 != nil {
		if p1.Val != val {
			p2.Next = &ListNode{Val: p1.Val}
			p2 = p2.Next
		}
		p1 = p1.Next
	}

	return out.Next
}

func RemoveElements2(head *ListNode, val int) *ListNode {
	return removeElements(head, val)
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} else if head.Val != val {
		head.Next = removeElements(head.Next, val)
		return head
	} else {
		head = removeElements(head.Next, val)
		return head
	}
}
