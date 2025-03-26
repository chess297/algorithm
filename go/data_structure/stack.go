package data_structure

type Stack struct {
	data []int
}

func NewStack(data []int) *Stack {
	return &Stack{data: data}
}

func (a *Stack) Push(index int) {
	a.data = append(a.data, index)
}

func (a *Stack) Pop() int {
	last := a.data[len(a.data)-1]
	a.data = a.data[:len(a.data)-1]
	return last
}

func (a *Stack) Size() int {
	return len(a.data)
}

func (a *Stack) IsEmpty() bool {
	return len(a.data) == 0
}
