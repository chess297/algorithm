package learn_go

import (
	"testing"
)

var (
	global_go_name = "GlobalGolang"
)

func TestVariable(t *testing.T) {
	// 变量一般使用:= 声明+赋值
	_go_name := "Golang"
	var go_name string // 如果只定义变量，没有赋值，那么变量的值就是该类型的零值
	// var 和 := 都是用来声明变量的，但是:= 只能用于函数内部，不能用于全局变量
	// var 可以用于全局变量，但是不能用于函数内部
	// var 可以用于函数内部，也可以用于全局变量

	t.Run("var 只定义不赋值，那么变量的值就是该类型的零值 \"\"", func(t *testing.T) {
		if go_name != "" {
			t.Errorf("go_name = %v, want %v", go_name, "")
		}
	})

	if true {
		go_name = "Golang"
	}

	t.Run("var 和:= 其实是等价的", func(t *testing.T) {
		if _go_name != go_name {
			t.Errorf("_go_name = %v, want %v", go_name, go_name)
		}
	})

	t.Run("global_go_name", func(t *testing.T) {
		if global_go_name != "GlobalGolang" {
			t.Errorf("go_name = %v, want %v", global_go_name, "Golang")
		}
	})
}
