package stack

import (
	"go-algorithm/data_structure/array"
)

type SliceArrayStack struct {
	Stack
	data array.SliceArray
}

func NewSliceArrayStack(data []int) *SliceArrayStack {
	return &SliceArrayStack{data: array.NewSliceArray(data)}
}

func (a *SliceArrayStack) Peek() int {
	return a.data.GetLast()
}
func (a *SliceArrayStack) Push(e int) {
	a.data.AddLast(e)
}

func (a *SliceArrayStack) Pop() int {
	if a.IsEmpty() {
		panic("SliceArray Dequeue Fail , SliceArray is empty")
	}
	top := a.Peek()
	a.data.DeleteLast()
	return top
}

func (a *SliceArrayStack) GetSize() int {
	return a.data.GetSize()
}

func (a *SliceArrayStack) IsEmpty() bool {
	return a.GetSize() == 0
}
