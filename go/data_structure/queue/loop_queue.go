package queue

// 循环队列
type LoopQueue struct {
	// Queue
	data  []int
	front int
	tail  int
}

func NewLoopQueue() LoopQueue {
	return LoopQueue{}
}

func (q *LoopQueue) Push(e int) {
	q.data = append(q.data, e)
}

func (q *LoopQueue) GetFront() int {
	return q.front
}

func (q *LoopQueue) GetTail() int {
	return q.tail
}
