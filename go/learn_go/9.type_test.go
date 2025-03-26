package learn_go

import "testing"

type Number interface {
	comparable
	int | float64 | string
}

func Sum[T Number](a, b T) T {
	return a + b
}

func TestSum(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if Sum(1, 2) != 3 {
			t.Errorf("Sum(1, 2) = %v, want %v", Sum(1, 2), 3)
		}
	})
}
