package listNode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	out := make([]int, 0)

	node := ln
	for node != nil {
		out = append(out, node.Val)
		node = node.Next
	}

	return fmt.Sprintf("%v", out)
}

func GenerateListNode(values ...int) *ListNode {

	var head, node *ListNode

	for _, v := range values {
		if node == nil {
			node = &ListNode{Val: v}
			head = node
			continue
		}
		node.Next = &ListNode{Val: v}
		node = node.Next
	}

	return head
}

func GetElement(head *ListNode, index int) *ListNode {

	node := head
	n := 0
	for node != nil {
		if n == index {
			return node
		}
		n += 1
		node = node.Next
	}

	return node
}

func Insert(head *ListNode, index, val int) *ListNode {
	if index == 0 {
		head = &ListNode{Val: val, Next: head}
		return head
	}

	node := head
	n := 0
	for node != nil {
		// 定位到前一个节点
		if n == index-1 {
			node.Next = &ListNode{Val: val, Next: node.Next}
			break
		}
		n += 1
		node = node.Next
	}

	return head
}

type DoubleNode struct {
	Val  int
	Prev *DoubleNode
	Next *DoubleNode
}

func DoubleNodeInsert(node *DoubleNode, val int) {

	next := node.Next
	cur := &DoubleNode{Val: val}
	cur.Next = next
	cur.Prev = node
	next.Prev = node
	node.Next = cur

}

func DoubleNodeRemove(node *DoubleNode) {

	prior := node.Prev
	next := node.Next
	prior.Next = next
	next.Prev = prior

	node = nil
}