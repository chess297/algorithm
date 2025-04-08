package l2_test

import "testing"

func TestArray(t *testing.T) {
	var arr [5]int // 数组的长度是固定的，不能改变 [] 不给定长度，就是切片
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	// arr[5] = 6 // 数组的长度是固定的，不能改变
	t.Run("数组的长度是固定的，不能改变", func(t *testing.T) {
		if arr[0] != 1 {
			t.Errorf("arr[0] = %v, want %v", arr[0], 1)
		}
		if len(arr) != 5 {
			t.Errorf("len(arr) = %v, want %v", len(arr), 5)
		}
	})

	// 数组的遍历
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

}
