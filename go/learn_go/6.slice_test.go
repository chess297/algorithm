package learn_go

import "testing"

// 切片的长度
func TestSliceLen(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Run("切片的长度", func(t *testing.T) {
		if len(slice) != 5 {
			t.Errorf("len(slice) = %v, want %v", len(slice), 5)
		}
	})
	t.Run("切片的值", func(t *testing.T) {
		if slice[0] != 1 {
			t.Errorf("slice[0] = %v, want %v", 1, 1)
		}
	})

	slice1 := make([]int, 5, 10) // make([]类型, 长度, 容量)
	t.Run("make 创建的切片长度", func(t *testing.T) {
		if len(slice1) != 5 {
			t.Errorf("len(slice1) = %v, want %v", len(slice1), 5)
		}
		if cap(slice1) != 10 {
			t.Errorf("cap(slice1) = %v, want %v", cap(slice1), 10)
		}
	})
	t.Run("make 创建的切片默认值为 对应类型的零值", func(t *testing.T) {
		if slice1[0] != 0 {
			t.Errorf("slice1[0] = %v, want %v", slice1[0], 0)
		}
	})

}

// 切片的容量
func TestSliceCap(t *testing.T) {
	slice := make([]int, 5, 10) // make([]类型, 长度, 容量)

	t.Run("切片的容量", func(t *testing.T) {
		if cap(slice) != 10 {
			t.Errorf("cap(slice) = %v, want %v", cap(slice), 10)
		}
	})

	t.Run("切片的默认值", func(t *testing.T) {
		if slice[0] != 0 {
			t.Errorf("slice[0] = %v, want %v", slice[0], 0)
		}
	})
	slice[0] = 1
	t.Run("切片赋值", func(t *testing.T) {
		if slice[0] != 1 {
			t.Errorf("slice[0] = %v, want %v", slice[0], 1)
		}
	})
	slice = append(slice, 1)
	t.Run("append 先增加的是长度", func(t *testing.T) {
		if cap(slice) != 10 {
			t.Errorf("cap(slice) = %v, want %v", cap(slice), 10)
		}
		if len(slice) != 6 {
			t.Errorf("len(slice) = %v, want %v", len(slice), 6)
		}
	})
	slice = append(slice, 1, 2, 3, 4, 5)
	t.Run("超过长度，会自动扩容", func(t *testing.T) {
		if cap(slice) != 20 {
			t.Errorf("cap(slice) = %v, want %v", cap(slice), 20)
		}
	})
}
