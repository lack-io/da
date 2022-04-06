package day02

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 三组映射关系:
	//  * 源链表元素 -> random 指向元素
	//  * 源链表元素 -> 复制链表元素
	//  * 复制链表未指向的元素
	srcMap := make(map[*Node]*Node, 0)
	dstMap := make(map[*Node]*Node, 0)
	unmapping := make([]*Node, 0)

	var out *Node
	var p1, p2 *Node = head, nil
	for p1 != nil {
		node := &Node{Val: p1.Val}
		if p1.Random != nil {
			v, ok := dstMap[srcMap[p1]]
			if ok {
				node.Random = v
			} else {
				unmapping = append(unmapping, p1)
			}

			srcMap[p1] = p1.Random
		}
		dstMap[p1] = node

		if p2 == nil {
			out = node
			p2 = node
		} else {
			p2.Next = node
			p2 = p2.Next
		}

		p1 = p1.Next
	}
	for _, node := range unmapping {
		dstMap[node].Random = dstMap[srcMap[node]]
	}

	return out
}
