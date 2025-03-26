package data_structure

type Stack[T any] struct {
	data []T
}

func NewStack[T any](data []T) *Stack[T] {
	return &Stack[T]{data: data}
}

func (a *Stack[T]) Top() T {
	return a.data[len(a.data)-1]
}

func (a *Stack[T]) Push(data T) {
	a.data = append(a.data, data)
}

func (a *Stack[T]) Pop() T {
	last := a.data[len(a.data)-1]
	a.data = a.data[:len(a.data)-1]
	return last
}

func (a *Stack[T]) Size() int {
	return len(a.data)
}

func (a *Stack[T]) IsEmpty() bool {
	return len(a.data) == 0
}
