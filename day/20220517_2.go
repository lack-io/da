package day

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return mergeList(head)
}

func mergeList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	var p, q, pre *ListNode = head, head, nil
	for q != nil && q.Next != nil {
		pre = p
		p = p.Next
		q = q.Next.Next
	}
	pre.Next = nil
	l := mergeList(head)
	r := mergeList(p)
	return merge(l, r)
}

func merge(l, r *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			cur = cur.Next
			l = l.Next
		} else {
			cur.Next = r
			cur = cur.Next
			r = r.Next
		}
	}
	if l != nil {
		cur.Next = l
	}
	if r != nil {
		cur.Next = r
	}
	return head.Next
}
