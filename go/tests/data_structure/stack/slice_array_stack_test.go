package stack_test

import (
	"go-algorithm/data_structure/stack"
	"testing"
)

func TestStack(t *testing.T) {
	stack := stack.NewSliceArrayStack([]int{})
	stack.Push(1)
	t.Run("SliceArrayStack Peek", func(t *testing.T) {
		if stack.Peek() != 1 {
			t.Errorf("expected 1, got %d", 2)
		}
	})
	t.Run("SliceArrayStack Push", func(t *testing.T) {
		stack.Push(2)
		if stack.Peek() != 2 {
			t.Errorf("expected 2, got %d", stack.Peek())
		}
	})
	last := stack.Pop()
	t.Run("SliceArrayStack Pop", func(t *testing.T) {
		if last != 2 {
			t.Errorf("expected 2, got %d", 2)
		}
		if stack.Peek() != 1 {
			t.Errorf("expected 2, got %d", 1)
		}
	})

}
