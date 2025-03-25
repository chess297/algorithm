package data_structure

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Append(1)
	stack.Append(2)
	t.Run("Append", func(t *testing.T) {
		if stack.Len() != 2 {
			t.Errorf("expected 2, got %d", stack.Len())
		}
	})
}
