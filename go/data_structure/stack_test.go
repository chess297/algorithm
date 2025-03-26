package data_structure

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack([]int{})
	stack.Push(1)
	stack.Push(2)
	t.Run("Append", func(t *testing.T) {
		if stack.Size() != 2 {
			t.Errorf("expected 2, got %d", stack.Size())
		}
	})
}
