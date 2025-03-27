package data_structure

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewArrayQueue(NewArray([]int{}, 10))
	queue.Enqueue(1)
	queue.Enqueue(2)
	t.Run("Front", func(t *testing.T) {
		e, ok := queue.GetFront()
		if e != 1 || !ok {
			t.Errorf("expected 1, got %d", 1)
		}
	})

	t.Run("Enqueue Size", func(t *testing.T) {
		if queue.GetSize() != 2 {
			t.Errorf("expected 2, got %d", queue.GetSize())
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		first, _ := queue.Dequeue()
		if first != 1 {
			t.Errorf("expected 1, got %d", 1)
		}
		if e, ok := queue.GetFront(); e != 2 || !ok {
			t.Errorf("expected 2, got %d", 2)
		}
		if queue.GetSize() != 1 {
			t.Errorf("expected 1, got %d", queue.GetSize())
		}
	})

}
