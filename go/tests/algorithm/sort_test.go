package algorithm_test

import (
	"go-algorithm/algorithm/sort"
	"testing"
)

func TestSort(t *testing.T) {
	list := []int{5, 4, 3, 2, 1}
	sort.SelectionSort(list)
	t.Run("SelectionSort", func(t *testing.T) {
		if list[0] != 1 || list[1] != 2 || list[2] != 3 || list[3] != 4 || list[4] != 5 {
			t.Errorf("SelectionSort() = %v, want %v", list[0], 1)
		}
	})

	list2 := []int{5, 3, 4, 2, 1}
	sort.InsertionSort(list2)
	t.Run("InsertionSort", func(t *testing.T) {
		if list[0] != 1 || list[1] != 2 || list[2] != 3 || list[3] != 4 || list[4] != 5 {
			t.Errorf("InsertionSort() = %v, want %v", list[0], 1)
		}
	})
}
