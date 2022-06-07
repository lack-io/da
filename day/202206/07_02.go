package _02206

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode = head, nil

	for fast != nil {
		if n == 0 {
			if slow == nil {
				slow = head
			} else {
				slow = slow.Next
			}
		}
		fast = fast.Next
		if n > 0 {
			n -= 1
		}
	}

	if slow == nil {
		return head.Next
	}
	slow.Next = slow.Next.Next
	return head
}
