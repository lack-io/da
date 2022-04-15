package day11

func getKthFromEnd(head *ListNode, k int) *ListNode {

	slow, fast := head, head

	for fast != nil {
		if k > 0 {
			k -= 1
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}

	return slow
}
