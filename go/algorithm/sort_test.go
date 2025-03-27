package algorithm

import (
	"testing"
)

func TestSort(t *testing.T) {
	list := []int{5, 4, 3, 2, 1}
	SelectSort(list)
	t.Run("SelectSort", func(t *testing.T) {
		if list[0] != 1 || list[1] != 2 || list[2] != 3 || list[3] != 4 || list[4] != 5 {
			t.Errorf("SelectSort() = %v, want %v", list[0], 1)
		}
	})

	list2 := []int{5, 4, 3, 2, 1}
	InsertSort(list2)
	t.Run("InsertSort", func(t *testing.T) {
		if list[0] != 1 || list[1] != 2 || list[2] != 3 || list[3] != 4 || list[4] != 5 {
			t.Errorf("InsertSort() = %v, want %v", list[0], 1)
		}
	})
}
