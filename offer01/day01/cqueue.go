package day01

type CQueue struct {
	head []int
	tail []int
}

func Constructor() CQueue {
	return CQueue{
		head: make([]int, 0),
		tail: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.tail = append(this.tail, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.head) > 0 {
		v := this.head[len(this.head)-1]
		this.head = this.head[:len(this.head)-1]
		return v
	}
	if len(this.tail) > 0 {
		for i := len(this.tail) - 1; i >= 0; i-- {
			this.head = append(this.head, this.tail[i])
		}
		this.tail = make([]int, 0)
		return this.DeleteHead()
	}

	return -1
}
