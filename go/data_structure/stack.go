package data_structure

type Stack struct {
	items []int
}

func NewStack() Stack {
	return Stack{
		items: []int{1, 2, 3, 4, 5},
	}
}

func (s *Stack) Pop() int {
	lastIndex := s.Len() - 1
	last := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return last
}

func (s *Stack) Len() int {
	return len(s.items)
}
