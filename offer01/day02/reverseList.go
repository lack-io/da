package day02

func reverseList(head *ListNode) *ListNode {
	var slow, fast *ListNode = nil, head

	for fast != nil {
		next := fast.Next
		fast.Next = slow
		slow = fast
		fast = next
	}

	return slow
}