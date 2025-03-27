package data_structure

type Stack[T any] interface {
	Push(e T)
	Pop() T
	Peek() T
	GetSize() int
	IsEmpty() bool
}

type ArrayStack[T any] struct {
	Stack[T]
	data Array[T]
}

func NewArrayStack[T any](capacity int) *ArrayStack[T] {
	return &ArrayStack[T]{data: Array[T]{
		data:     []T{},
		capacity: capacity,
		size:     0,
	}}
}

func (a *ArrayStack[T]) Peek() (T, bool) {
	return a.data.GetLast()
}
func (a *ArrayStack[T]) Push(e T) {
	a.data.Add(e)
}

func (a *ArrayStack[T]) Pop() (T, bool) {
	top, ok := a.Peek()

	if !ok {
		var zero T
		return zero, false
	}
	a.data.DeleteLast()
	return top, true
}

func (a *ArrayStack[T]) GetSize() int {
	return a.data.GetSize()
}

func (a *ArrayStack[T]) IsEmpty() bool {
	return a.data.size == 0
}
