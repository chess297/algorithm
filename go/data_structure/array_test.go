package data_structure

import "testing"

func TestArray(t *testing.T) {
	a := &Array[int]{
		data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	t.Run("Array size", func(t *testing.T) {
		if a.GetSize() != 9 {
			t.Errorf("Array size is %d want %d", a.GetSize(), 9)
		}
	})

	t.Run("Array size", func(t *testing.T) {
		a := &Array[int]{data: []int{0, 1}}
		e, _ := a.Get(1)
		if e != 1 {
			t.Errorf("Array Find is %v want %d", e, 1)
		}
	})

	t.Run("Array Delete", func(t *testing.T) {
		a := &Array[int]{data: []int{1}}
		a.Delete(0)
		if a.GetSize() != 0 {
			t.Errorf("Array GetSize  want %v", nil)
		}
	})
}
