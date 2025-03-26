package data_structure

type Queue[T comparable] struct {
	data []T
}

func NewQueue[T comparable](data []T) *Queue[T] {
	return &Queue[T]{data: data}
}

func (a *Queue[T]) Enqueue(data T) {
	a.data = append(a.data, data)
}

func (a *Queue[T]) Dequeue() T {
	first := a.data[0]
	a.data = a.data[1:]
	return first
}

func (a *Queue[T]) Size() int {
	return len(a.data)
}

func (a *Queue[T]) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *Queue[T]) Front() T {
	return a.data[0]
}

func (a *Queue[T]) Back() T {
	return a.data[len(a.data)-1]
}

func (a *Queue[T]) Clear() {
	a.data = []T{}
}

// func (a *Queue[T]) Contains(data T) bool {
// 	for _, v := range a.data {
// 		if equal() v == data {
// 			return true
// 		}
// 	}
// 	return false
// }
