package data_structure

type Queue interface {
}

type ArrayQueue[T comparable] struct {
	Queue
	data []T
}

func NewArrayQueue[T comparable](data []T) *ArrayQueue[T] {
	return &ArrayQueue[T]{data: data}
}

func (a *ArrayQueue[T]) Enqueue(data T) {
	a.data = append(a.data, data)
}

func (a *ArrayQueue[T]) Dequeue() T {
	first := a.data[0]
	a.data = a.data[1:]
	return first
}

func (a *ArrayQueue[T]) Size() int {
	return len(a.data)
}

func (a *ArrayQueue[T]) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *ArrayQueue[T]) Front() T {
	return a.data[0]
}

func (a *ArrayQueue[T]) Back() T {
	return a.data[len(a.data)-1]
}

func (a *ArrayQueue[T]) Clear() {
	a.data = []T{}
}

// func (a *ArrayQueue[T]) Contains(data T) bool {
// 	for _, v := range a.data {
// 		if equal() v == data {
// 			return true
// 		}
// 	}
// 	return false
// }
