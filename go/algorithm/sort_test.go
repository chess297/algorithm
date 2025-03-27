package algorithm

import (
	"testing"
)

func TestMain(t *testing.T) {
	list := []int{5, 4, 3, 2, 1}
	SelectSort(list)
	t.Run("min sort", func(t *testing.T) {
		if list[0] != 1 || list[1] != 2 || list[2] != 3 || list[3] != 4 || list[4] != 5 {
			t.Errorf("MinSort() = %v, want %v", list[0], 1)
		}
	})
}
