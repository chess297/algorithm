package data_structure

type Stack struct {
	items []int
}

func NewStack() Stack {
	return Stack{
		items: []int{},
	}
}

func (s *Stack) Pop() int {
	lastIndex := s.Len() - 1
	last := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return last
}

func (s *Stack) Append(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Len() int {
	return len(s.items)
}
