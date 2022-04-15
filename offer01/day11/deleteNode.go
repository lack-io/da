package day11

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {

	fast := head
	var slow *ListNode

	for fast != nil {
		if fast.Val == val {
			if slow == nil {
				head = fast.Next
			} else {
				slow.Next = fast.Next
			}

			break
		}

		slow = fast
		fast = fast.Next
	}

	return head
}
