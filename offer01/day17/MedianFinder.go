package day17

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

type MedianFinder struct {
	maxHeap, minHeap *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	a := &IntHeap{}
	b := &IntHeap{}
	heap.Init(a)
	heap.Init(b)

	return MedianFinder{a, b}

}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num > (*this.maxHeap)[0] {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, -num)
	}

	// 调节大小栈，确保 `FindMedian()` 得到正确的结果
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
	} else if this.minHeap.Len() > this.maxHeap.Len()+1 {
		heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() < this.minHeap.Len() {
		return float64(-(*this.minHeap)[0])
	} else if this.maxHeap.Len() > this.minHeap.Len() {
		return float64((*this.maxHeap)[0])
	}

	return float64(-(*this.minHeap)[0]+(*this.maxHeap)[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
