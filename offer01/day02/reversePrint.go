package day02

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint1(head *ListNode) []int {
	out := make([]int, 0)

	for node := head; node != nil; node = node.Next {
		out = append([]int{node.Val}, out...)
	}

	return out
}

func reversePrint2(head *ListNode) []int {
	out := make([]int, 0)
	stack := make([]*ListNode, 0)

	for node := head; node != nil; node = node.Next {
		stack = append(stack, node)
	}

	for i := len(stack) - 1; i >= 0; i++ {
		out = append(out, stack[i].Val)
	}

	return out
}
