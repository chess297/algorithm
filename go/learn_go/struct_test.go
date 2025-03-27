package learn_go

import (
	"go-algorithm/data_structure"
	"testing"
)

// 测试栈的结构体 更多结构体的特性，可以看 data_structure/stack_test.go
func TestStackStruct(t *testing.T) {
	stack := data_structure.NewArray([]int{1, 2, 3, 4, 5}, 10)
	// stack.data // 不能访问，因为 data 是 private 的
	stack.Add(6)
	t.Run("往栈内添加元素", func(t *testing.T) {
		if stack.GetSize() != 6 {
			t.Errorf("stack.data[0] = %v, want %v", stack.GetSize(), 6)
		}
	})

}
