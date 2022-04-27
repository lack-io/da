package heap

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxHeap struct {
	data *IntHeap
}

func NewMaxHeap() *MaxHeap {
	data := &IntHeap{}
	heap.Init(data)
	return &MaxHeap{data: data}
}

func (h *MaxHeap) Push(v int) {
	heap.Push(h.data, v)
}

func (h *MaxHeap) Pop() (int, error) {
	if h.data.Len() == 0 {
		return -1, fmt.Errorf("empty heap")
	}
	v := heap.Pop(h.data)
	return v.(int), nil
}

func (h *MaxHeap) Top() (int, error) {
	if h.data.Len() == 0 {
		return -1, fmt.Errorf("empty heap")
	}
	return (*h.data)[0], nil
}
