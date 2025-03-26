package data_structure

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue([]int{})
	queue.Enqueue(1)
	queue.Enqueue(2)
	t.Run("Front", func(t *testing.T) {
		if queue.Front() != 1 {
			t.Errorf("expected 1, got %d", 1)
		}
	})

	t.Run("Enqueue Size", func(t *testing.T) {
		if queue.Size() != 2 {
			t.Errorf("expected 2, got %d", queue.Size())
		}
	})

	first := queue.Dequeue()
	t.Run("Dequeue", func(t *testing.T) {
		if first != 1 {
			t.Errorf("expected 1, got %d", 1)
		}
		if queue.Front() != 2 {
			t.Errorf("expected 2, got %d", 2)
		}
	})
}
