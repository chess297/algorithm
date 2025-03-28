package stack

type Stack interface {
	Push(e int)
	Pop() int
	Peek() int
	GetSize() int
	IsEmpty() bool
}
