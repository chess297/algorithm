package queue_test

import (
	"go-algorithm/data_structure/queue"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := queue.NewSliceArrayQueue([]int{})
	queue.Enqueue(1)
	queue.Enqueue(2)
	t.Run("Front", func(t *testing.T) {
		if queue.GetFront() != 1 {
			t.Errorf("expected 1, got %d", 1)
		}
	})

	t.Run("Enqueue Size", func(t *testing.T) {
		if queue.GetSize() != 2 {
			t.Errorf("expected 2, got %d", queue.GetSize())
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		first := queue.Dequeue()
		if first != 1 {
			t.Errorf("expected 1, got %d", 1)
		}
		if queue.GetFront() != 2 {
			t.Errorf("expected 2, got %d", 2)
		}
		if queue.GetSize() != 1 {
			t.Errorf("expected 1, got %d", queue.GetSize())
		}
	})

}
