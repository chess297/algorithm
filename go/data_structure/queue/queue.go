package queue

type Queue interface {
	Enqueue(e int)
	Dequeue() int
	GetFront() int
	GetSize() int
	IsEmpty() bool
}
