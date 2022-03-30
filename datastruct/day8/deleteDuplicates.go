package day8

func DeleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head.Next
	for p2 != nil {
		if p1.Val == p2.Val {
			p1.Next = p2.Next
		} else {
			p1 = p1.Next
		}
		p2 = p2.Next
	}

	return head
}

func DeleteDuplicates2(head *ListNode) *ListNode {
	return deleteDuplicates(head)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		head.Next = deleteDuplicates(head.Next.Next)
		return deleteDuplicates(head)
	}
	deleteDuplicates(head.Next.Next)
	return head
}
