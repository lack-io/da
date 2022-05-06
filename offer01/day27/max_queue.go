package day27

import "container/list"

type MaxQueue struct {
	queue *list.List
	deque *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: list.New(),
		deque: list.New(),
	}
}

func (q *MaxQueue) Max_value() int {
	if q.deque.Len() == 0 {
		return -1
	}
	return q.deque.Front().Value.(int)
}

func (q *MaxQueue) Push_back(value int) {
	q.queue.PushBack(value)
	for q.deque.Len() > 0 && q.deque.Back().Value.(int) < value {
		q.deque.Remove(q.deque.Back())
	}
	q.deque.PushBack(value)
}

func (q *MaxQueue) Pop_front() int {
	if q.queue.Len() == 0 {
		return -1
	}
	if q.queue.Front().Value.(int) == q.deque.Front().Value.(int) {
		q.deque.Remove(q.deque.Front())
	}
	return q.queue.Remove(q.queue.Front()).(int)
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
