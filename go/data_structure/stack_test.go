package data_structure

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewArrayStack[int](10)
	stack.Push(1)
	stack.Push(2)
	// t.Run("Top", func(t *testing.T) {
	// 	if stack.Peek() != 2 {
	// 		t.Errorf("expected 2, got %d", 2)
	// 	}
	// })
	// t.Run("Push", func(t *testing.T) {
	// 	if stack.GetSize() != 2 {
	// 		t.Errorf("expected 2, got %d", stack.GetSize())
	// 	}
	// })
	// last := stack.Pop()
	// t.Run("Pop", func(t *testing.T) {
	// 	if last != 2 {
	// 		t.Errorf("expected 2, got %d", 2)
	// 	}
	// 	if stack.Peek() != 1 {
	// 		t.Errorf("expected 2, got %d", 1)
	// 	}
	// })

}
