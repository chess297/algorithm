package array_test

import (
	"go-algorithm/data_structure/array"
	"go-algorithm/utils"
	"testing"
)

func TestArray(t *testing.T) {
	a := array.NewSliceArray(utils.GenOrderIntList(10))
	t.Run("SliceArray GetSize", func(t *testing.T) {
		if a.GetSize() != 10 {
			t.Errorf("expected %d want %d", 10, a.GetSize())
		}
	})

	t.Run("SliceArray Get", func(t *testing.T) {
		if a.Get(0) != 1 {
			t.Errorf("expected %v want %d", 1, a.Get(0))
		}
	})

	t.Run("SliceArray Delete", func(t *testing.T) {
		a.Delete(0)
		if a.GetSize() != 9 {
			t.Errorf("expected %d want %d", 9, a.GetSize())
		}
	})
}
