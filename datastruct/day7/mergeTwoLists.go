package day7

func MergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2

	cur := &ListNode{}
	head := cur
	for {
		if p1 == nil {
			cur.Next = p2
			break
		}
		if p2 == nil {
			cur.Next = p1
			break
		}

		if p1.Val < p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}

		cur = cur.Next
	}

	return head.Next
}

func MergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	if p1 == nil {
		return p2
	} else if p2 == nil {
		return p1
	} else if p1.Val < p2.Val {
		p1.Next = MergeTwoLists2(p1.Next, p2)
		return p1
	} else {
		p2.Next = MergeTwoLists2(p1, p2.Next)
		return p2
	}
}
