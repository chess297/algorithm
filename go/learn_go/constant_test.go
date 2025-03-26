package learn_go

import "testing"

const (
	GO_NAME = "Golang"
)

func TestConstant(t *testing.T) {
	// GO_NAME = "Golang" // 常量不可修改
	t.Run("const 常量", func(t *testing.T) {
		if GO_NAME != "Golang" {
			t.Errorf("GO_NAME = %v, want %v", GO_NAME, "Golang")
		}
	})
}
