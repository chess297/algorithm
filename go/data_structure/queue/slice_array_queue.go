package queue

import (
	"go-algorithm/data_structure/array"
)

type SliceArrayQueue struct {
	Queue
	data array.SliceArray
}

func NewSliceArrayQueue(data []int) *SliceArrayQueue {
	return &SliceArrayQueue{data: array.NewSliceArray(data)}
}

func (a *SliceArrayQueue) Enqueue(e int) {
	a.data.AddLast(e)
}

func (a *SliceArrayQueue) Dequeue() int {
	if a.IsEmpty() {
		panic("SliceArray Dequeue Fail , SliceArray is empty")
	}
	e := a.data.GetFirst()
	a.data.DeleteFirst()
	return e
}

func (a *SliceArrayQueue) GetFront() int {
	return a.data.GetFirst()
}

func (a *SliceArrayQueue) GetSize() int {
	return a.data.GetSize()
}

func (a *SliceArrayQueue) IsEmpty() bool {
	return a.data.IsEmpty()
}
