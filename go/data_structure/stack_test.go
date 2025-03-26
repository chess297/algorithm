package data_structure

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack([]int{})
	stack.Push(1)
	stack.Push(2)
	t.Run("Top", func(t *testing.T) {
		if stack.Top() != 2 {
			t.Errorf("expected 2, got %d", 2)
		}
	})
	t.Run("Push", func(t *testing.T) {
		if stack.Size() != 2 {
			t.Errorf("expected 2, got %d", stack.Size())
		}
		if len(stack.data) != 2 { // 因为在同一个包内，所以data能访问到
			t.Errorf("expected 2, got %d", stack.Size())
		}
	})
	last := stack.Pop()
	t.Run("Pop", func(t *testing.T) {
		if last != 2 {
			t.Errorf("expected 2, got %d", 2)
		}
		if stack.Top() != 1 {
			t.Errorf("expected 2, got %d", 1)
		}
	})

}
