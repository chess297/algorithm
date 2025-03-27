package data_structure

type Queue[T any] interface {
	Enqueue(e T)
	Dequeue() (T, bool)
	GetFront() (T, bool)
	GetSize() int
	IsEmpty() bool
}

type ArrayQueue[T comparable] struct {
	Queue[T]
	data Array[T]
}

func NewArrayQueue[T comparable](data Array[T]) *ArrayQueue[T] {
	return &ArrayQueue[T]{data: data}
}

func (a *ArrayQueue[T]) Enqueue(data T) {
	// a.data = append(a.data, data)
	a.data.AddLast(data)
}

func (a *ArrayQueue[T]) Dequeue() (T, bool) {
	first, ok := a.GetFront()
	if !ok {
		var zero T
		return zero, false
	}
	a.data.DeleteFirst()
	return first, true
}

func (a *ArrayQueue[T]) GetFront() (T, bool) {
	return a.data.GetFirst()
}

func (a *ArrayQueue[T]) GetSize() int {
	return a.data.GetSize()
}

func (a *ArrayQueue[T]) IsEmpty() bool {
	return a.GetSize() == 0
}
